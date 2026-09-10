package controller

import "github.com/mrandrey24/url-shortening-server/internal/domain"

func domainURLToResponseCode(url *domain.URL) *URLResponseGetByCode {
	return &URLResponseGetByCode{
		ID:        url.ID,
		URL:       url.URL,
		ShortCode: url.ShortCode,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}
}

func domainURLToResponseStats(url *domain.URL) *URLResponseGetStats {
	return &URLResponseGetStats{

		ID:          url.ID,
		URL:         url.URL,
		ShortCode:   url.ShortCode,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		AccessCount: url.AccessCount,
	}
}

func domainURLToResponseUpdate(url *domain.URL) *URLResponseUpdate {
	return &URLResponseUpdate{
		URL: url.URL,
	}
}
