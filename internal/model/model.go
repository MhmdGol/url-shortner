package model

import "time"

type ID int64

type URLModel struct {
	ID           ID
	URLDomain    string
	HashString   string
	LongURL      string
	ClickedCount int64
	CreatedBy    ID
	CreatedAt    time.Time
}
