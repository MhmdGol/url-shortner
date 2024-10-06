package service

import (
	"context"
	"time"
	"url-shortner/internal/domain"
)

type URLShortnerService struct{}

func (u *URLShortnerService) ShortURL(ctx context.Context, url string) string {
	shortURL := domain.ShortURL{
		URLDomain: "moh.com",
		LongURL:   url,
		CreatedAt: time.Now(),
	}

	shortURL.GenerateShortURL()

	return shortURL.URLDomain + "/" + shortURL.HashString
}
