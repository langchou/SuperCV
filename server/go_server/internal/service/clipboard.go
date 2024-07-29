package service

import (
	"context"
	"database/sql"
	"fmt"
	"go_server/internal/model"
	"go_server/pkg/logger"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func createClipTable() error {
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS clips (
            id SERIAL PRIMARY KEY,
            content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)

	if err != nil {
		logger.ErrorLogger.Printf("Failed to create clip table: %v", err)
		return err
	}
	return nil
}

func createUserTable() error {
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)

	if err != nil {
		logger.ErrorLogger.Printf("Failed to create users table: %v", err)
		return err
	}

	logger.InfoLogger.Printf("create users table successfully")
	return nil
}

func createDeviceTable() error {
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS devices (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			uuid VARCHAR(255) NOT NULL UNIQUE,
			user_name VARCHAR(255) NOT NULL,
			status BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_name) REFERENCES users(name)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create devices table: %w", err)
	}
	// 创建更新 updated_at 的触发器
	_, err = db.Exec(`
        CREATE OR REPLACE FUNCTION update_device_timestamp()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.updated_at = NOW();
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;

        DROP TRIGGER IF EXISTS update_device_timestamp ON devices;

        CREATE TRIGGER update_device_timestamp
        BEFORE UPDATE ON devices
        FOR EACH ROW
        EXECUTE FUNCTION update_device_timestamp();
    `)

	if err != nil {
		return fmt.Errorf("failed to create update_device_timestamp trigger: %w", err)
	}

	return nil
}

func init() {
	var err error
	// 替换以下连接字符串以匹配PostgreSQL配置
	db, err = sql.Open("postgres", "host=localhost port=5432 user=jonty dbname=supercv password=abcd1234 sslmode=disable")
	if err != nil {
		panic(err)
	}
	// 创建表
	err = createClipTable()
	if err != nil {
		logger.ErrorLogger.Printf("something wrong")
	}

	err = createUserTable()
	if err != nil {
		logger.ErrorLogger.Printf("something wrong")
	}

	err = createDeviceTable()
	if err != nil {
		logger.ErrorLogger.Printf("something wrong")
	}

}

func SaveClipboard(clip model.Clip) error {
	_, err := db.Exec("INSERT INTO clips (content) VALUES ($1)", clip.Content)
	return err
}

func SaveUser(ctx context.Context, user model.User) error {
	_, err := db.ExecContext(ctx, `
	INSERT INTO users (name, created_at)
    VALUES ($1, $2)
	`, user.Name, time.Now())

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// 23505 是 PostgreSQL 唯一约束违反的错误代码
			return fmt.Errorf("user with name %s already exists", user.Name)
		}
		logger.ErrorLogger.Printf("Failed to save user: %v", err)
		return err
	}
	logger.InfoLogger.Printf("New user saved: %s", user.Name)
	return nil
}

func SaveDevice(ctx context.Context, device model.Device) error {
	_, err := db.ExecContext(ctx, `
	INSERT INTO devices (name, uuid, user_name, status, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6)
	`, device.Name, device.Uuid, device.UserName, device.Status, time.Now(), time.Now())

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// 23505 是 PostgreSQL 唯一约束违反的错误代码
			return fmt.Errorf("Device with name %s already exists", device.Name)
		}
		logger.ErrorLogger.Printf("Failed to save device: %v", err)
		return err
	}
	logger.InfoLogger.Printf("New device saved: %s", device.Name)
	return nil
}

func PingDB() error {
	return db.Ping()
}
