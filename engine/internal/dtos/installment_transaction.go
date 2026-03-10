package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type InstallmentTransactionResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Value       float64             `json:"value"`
	InitialDate time.Time           `json:"initial_date"`
	FinalDate   time.Time           `json:"final_date"`
	Category    CategoryResponse    `json:"category"`
	Creditcard  *CreditCardResponse `json:"creditcard"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type ShortInstallmentTransactionResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InstallmentTransactionRelations struct {
	CategoryResponse   CategoryResponse
	CreditcardResponse *CreditCardResponse
}

type InstallmentTransactionRequest struct {
	Name         string     `json:"name" binding:"required"`
	Value        float64    `json:"value" binding:"required"`
	InitialDate  time.Time  `json:"initial_date" binding:"required"`
	FinalDate    time.Time  `json:"final_date" binding:"required"`
	CategoryID   uuid.UUID  `json:"category_id" binding:"required"`
	CreditCardID *uuid.UUID `json:"creditcard_id"`
}

func (m InstallmentTransactionRequest) Validate() []errors.ApiErrorItem {
	errs := make([]errors.ApiErrorItem, 0)

	if utils.IsBlank(m.Name) {
		errs = append(errs, errors.InvalidFieldError(constants.NameEmptyMsg))
	}

	if len(m.Name) < 2 || len(m.Name) > 255 {
		errs = append(errs, errors.InvalidFieldError(constants.NameInvalidCharsCountMsg))
	}

	if m.Value < 0 || m.Value >= 1000000000000000.00 {
		errs = append(errs, errors.InvalidFieldError(constants.ValueInvalidMsg))
	}

	if m.InitialDate.IsZero() {
		errs = append(errs, errors.InvalidFieldError(constants.InitialDateEmptyMsg))
	}

	if m.FinalDate.IsZero() {
		errs = append(errs, errors.InvalidFieldError(constants.FinalDateEmptyMsg))
	}

	if !m.InitialDate.IsZero() && !m.FinalDate.IsZero() && m.FinalDate.Before(m.InitialDate) {
		errs = append(errs, errors.InvalidFieldError(constants.FinalDateBeforeInitialDateMsg))
	}

	if m.InitialDate.Equal(m.FinalDate) {
		errs = append(errs, errors.InvalidFieldError(constants.InitialDateEqualsFinalDateMsg))
	}

	return errs
}
