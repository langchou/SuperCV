package model

import "time"

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	EncryptedDEK string    `json:"encrypted_dek"`
	CreatedAt    time.Time `json:"created_at"`
}
