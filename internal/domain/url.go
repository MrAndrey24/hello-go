package domain

import (
	"time"
)

type URL struct {
	ID          int64     `json:"id"`
	URL         string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessCount int64     `json:"accessCount"`
}

// Method
func (u *URL) IncrementAccessCount() {
	u.AccessCount++
	u.UpdatedAt = time.Now()
}

type URLRepository interface {
	Save(u *URL) error
	Update(shortCode string, u *URL) (*URL, error)
	Delete(id string) error
	FindByShortCode(shortCode string) (*URL, error)
	FindUrlStatsByShortCode(shortCode string) (*URL, error)
}
