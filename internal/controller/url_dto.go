package controller

import (
	"time"
)

type URLResponseGetByCode struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type URLResponseGetStats struct {
	ID          int64     `json:"id"`
	URL         string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessCount int64     `json:"accessCount"`
}

type URLResponseUpdate struct {
	URL string `json:"url"`
}
