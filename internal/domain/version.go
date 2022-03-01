package domain

import "time"

type Version struct {
	Id        string    `json:"id"`
	ServiceId string    `json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
