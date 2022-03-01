package domain

import "time"

type Service struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Versions     []string  `json:"versions,omitempty"`
	VersionCount int       `json:"version_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
