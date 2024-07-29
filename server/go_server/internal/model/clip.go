package model

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type Clip struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
