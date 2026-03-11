package dtos

import (
	"backend-commons/constants"
	"time"

	"github.com/TB-Systems/go-commons/errors"
	"github.com/TB-Systems/go-commons/utils"
	"github.com/google/uuid"
)

type MonthlyTransactionResponse struct {
	ID         uuid.UUID           `json:"id"`
	Name       string              `json:"name"`
	Value      float64             `json:"value"`
	Day        int32               `json:"day"`
	Category   CategoryResponse    `json:"category"`
	Creditcard *CreditCardResponse `json:"creditcard"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

type ShortMonthlyTransactionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Day       int32     `json:"day"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MonthlyTransactionRelations struct {
	CategoryResponse   CategoryResponse
	CreditcardResponse *CreditCardResponse
}

type MonthlyTransactionRequest struct {
	Name         string     `json:"name" binding:"required"`
	Value        float64    `json:"value" binding:"required"`
	Day          int32      `json:"day" binding:"required"`
	CategoryID   uuid.UUID  `json:"category_id" binding:"required"`
	CreditCardID *uuid.UUID `json:"creditcard_id"`
}

func (m MonthlyTransactionRequest) Validate() []errors.ApiErrorItem {
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

	if m.Day < 1 || m.Day > 31 {
		errs = append(errs, errors.InvalidFieldError(constants.DayInvalidMsg))
	}

	return errs
}
