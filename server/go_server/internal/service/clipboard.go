package service

import (
	"database/sql"
	"go_server/internal/model"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	// 替换以下连接字符串以匹配你的 PostgreSQL 配置
	db, err = sql.Open("postgres", "host=localhost port=5432 user=jonty dbname=superCV password=abcd1234 sslmode=disable")
	if err != nil {
		panic(err)
	}

	// 创建表
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS clips (
            id SERIAL PRIMARY KEY,
            content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		panic(err)
	}
}

func SaveClipboard(clip model.Clip) error {
	_, err := db.Exec("INSERT INTO clips (content) VALUES ($1)", clip.Content)
	return err
}

func PingDB() error {
	return db.Ping()
}
