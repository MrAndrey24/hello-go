package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/mrandrey24/url-shortening-server/internal/domain"
)

// Dependency injection
type URLService struct {
	repo domain.URLRepository
}

// Constructor for URLService
func NewURLService(repo domain.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) CreateShortUrl(req string) (*domain.URL, error) {
	shortCode := generateShortCode()

	newUrl := &domain.URL{
		URL:       req,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Save(newUrl); err != nil {
		return nil, err
	}

	return newUrl, nil
}

func (s *URLService) UpdateShortUrl(code string, req *domain.URL) (*domain.URL, error) {
	if code == "" || req == nil || req.URL == "" {
		return nil, errors.New("Invalid Data")
	}
	newShortUrl, err := s.repo.Update(code, req)

	if err != nil {
		return nil, err
	}

	return newShortUrl, nil
}

func (s *URLService) RetrieveByShortCode(code string) (*domain.URL, error) {
	return s.repo.FindByShortCode(code)
}

func (s *URLService) DeleteSHortUrl(code string) error {
	return s.repo.Delete(code)
}

func generateShortCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = byte('a' + rand.Intn(26))
	}
	return string(b)
}
