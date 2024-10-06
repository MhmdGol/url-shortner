package domain

import (
	"crypto/sha256"
	"time"
	"url-shortner/internal/model"
)

type ShortURL struct {
	ID           model.ID
	URLDomain    string
	HashString   string
	LongURL      string
	ClickedCount int64
	CreatedBy    model.ID
	CreatedAt    time.Time
}

func (s *ShortURL) GenerateShortURL() {
	hasher := sha256.New()
	hasher.Write([]byte(s.LongURL))
	s.HashString = string(hasher.Sum(nil)[:8])
}
