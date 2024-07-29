package model

import "time"

type Device struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Uuid       string    `json:"user_uuid"`
	DeviceUuid string    `json:"device_uuid"`
	UserName   string    `json:"user_name"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
