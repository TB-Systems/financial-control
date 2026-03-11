package models

import (
	"time"

	"github.com/google/uuid"
)

type InstallmentTransaction struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        string
	Value       float64
	InitialDate time.Time
	FinalDate   time.Time
	Category    Category
	Creditcard  *CreditCard
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ShortInstallmentTransaction struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Name         string
	InitialDate  time.Time
	FinalDate    time.Time
	Value        float64
	CategoryID   uuid.UUID
	CreditCardID *uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateInstallmentTransaction struct {
	UserID       uuid.UUID
	Name         string
	Value        float64
	InitialDate  time.Time
	FinalDate    time.Time
	CategoryID   uuid.UUID
	CreditCardID *uuid.UUID
}
