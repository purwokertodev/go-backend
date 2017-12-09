package model

import (
	"time"
)

type Member struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	PasswordSalt string
	BirthDate    time.Time
	Version      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewMember() *Member {
	now := time.Now()
	return &Member{
		Version:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
