package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                     uuid.UUID
	UserID                 uuid.UUID
	Name                   string
	Date                   time.Time
	Value                  float64
	Paid                   bool
	Category               Category
	Creditcard             *CreditCard
	MonthlyTransaction     *ShortMonthlyTransaction
	AnnualTransaction      *ShortAnnualTransaction
	InstallmentTransaction *ShortInstallmentTransaction
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type CreateTransaction struct {
	UserID                   uuid.UUID
	Name                     string
	Date                     time.Time
	Value                    float64
	Paid                     bool
	CategoryID               uuid.UUID
	CreditcardID             *uuid.UUID
	MonthlyTransactionID     *uuid.UUID
	AnnualTransactionID      *uuid.UUID
	InstallmentTransactionID *uuid.UUID
}

type ShortTransaction struct {
	ID        uuid.UUID
	Name      string
	Date      time.Time
	Value     float64
	Paid      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionsCreditCardTotal struct {
	Date         time.Time
	UserID       uuid.UUID
	CreditcardID uuid.UUID
}
