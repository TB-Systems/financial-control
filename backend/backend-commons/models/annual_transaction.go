package models

import (
	"time"

	"github.com/google/uuid"
)

type AnnualTransaction struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Name       string
	Value      float64
	Day        int32
	Month      int32
	Category   Category
	Creditcard *CreditCard
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ShortAnnualTransaction struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Name         string
	Day          int32
	Month        int32
	Value        float64
	CategoryID   uuid.UUID
	CreditCardID *uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateAnnualTransaction struct {
	UserID       uuid.UUID
	Name         string
	Value        float64
	Day          int32
	Month        int32
	CategoryID   uuid.UUID
	CreditCardID *uuid.UUID
}
