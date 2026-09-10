package local

import (
	"errors"
	"time"

	"github.com/mrandrey24/url-shortening-server/internal/domain"
)

// Dependency injection
type memoryRepository struct {
	urls   []*domain.URL
	nextID int64
}

// Constructor for memoryRepository
func NewMemoryURLRepository() domain.URLRepository {
	return &memoryRepository{
		urls:   make([]*domain.URL, 0),
		nextID: 1,
	}
}

func (r *memoryRepository) Save(req *domain.URL) error {
	req.ID = r.nextID
	r.nextID++
	r.urls = append(r.urls, req)
	return nil
}

func (r *memoryRepository) Update(code string, req *domain.URL) (*domain.URL, error) {
	url, err := r.FindUrlStatsByShortCode(code)

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
			return nil
		}
	}
	return errors.New("ShortCode not found")
}

func (r *memoryRepository) FindByShortCode(shortCode string) (*domain.URL, error) {
	for _, url := range r.urls {
		if url.ShortCode == shortCode {
			url.IncrementAccessCount()
			return url, nil
		}
	}
	return nil, errors.New("ShortCode not found")
}

func (r *memoryRepository) FindUrlStatsByShortCode(shortCode string) (*domain.URL, error) {
	for _, items := range r.urls {
		if items.ShortCode == shortCode {
			return items, nil
		}
	}

	return nil, errors.New("ShortCode not found")
}
