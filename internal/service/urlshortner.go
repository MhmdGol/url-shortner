package service

import (
	"context"
	"time"
	"url-shortner/internal/domain"
	"url-shortner/internal/repository"
)

type URLShortnerService struct {
	repo repository.URLShortnerRepository
}

func (u *URLShortnerService) ShortURL(ctx context.Context, url string) string {
	shortURL := domain.ShortURL{
		URLDomain: "moh.com",
		LongURL:   url,
		CreatedAt: time.Now(),
	}

	shortURL.GenerateShortURL()

	return shortURL.URLDomain + "/" + shortURL.HashString
}

func (u *URLShortnerService) GetOriginalURL(ctx context.Context, shortenURL string) (string, error) {
	url, err := u.repo.Get(shortenURL)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}
