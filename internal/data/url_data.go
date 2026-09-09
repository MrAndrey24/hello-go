package local

import (
	"errors"
	"time"

	"github.com/mrandrey24/url-shortening-server/internal/domain"
)

// Dependency injection
type memoryRepository struct {
	urls []*domain.URL
}

// Constructor for memoryRepository
func NewMemoryURLRepository() domain.URLRepository {
	return &memoryRepository{
		urls: make([]*domain.URL, 0),
	}
}

func (r *memoryRepository) Save(req *domain.URL) error {
	req.ID = int64(len(r.urls) + 1)
	r.urls = append(r.urls, req)
	return nil
}

func (r *memoryRepository) Update(code string, req *domain.URL) (*domain.URL, error) {
	url, err := r.existShortCode(code)

	if err != nil {
		return nil, errors.New("ShortCode not found")
	}

	url.URL = req.URL
	url.UpdatedAt = time.Now()

	return url, err
}

func (r *memoryRepository) Delete(id string) error {
	for i, url := range r.urls {
		if url.ShortCode == id {
			r.urls = append(r.urls[:i], r.urls[i+1:]...)
			break
		}
	}
	return nil
}

func (r *memoryRepository) FindByShortCode(shortCode string) (*domain.URL, error) {
	for _, url := range r.urls {
		if url.ShortCode == shortCode {
			url.IncrementAccessCount()
			return url, nil
		}
	}
	return nil, nil
}

func (r *memoryRepository) existShortCode(code string) (*domain.URL, error) {
	for _, items := range r.urls {
		if items.ShortCode == code {
			return items, nil
		}
	}

	return nil, errors.New("ShortCode not found")
}
