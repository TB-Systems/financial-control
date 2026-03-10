package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionRelations struct {
	CategoryModel      models.Category
	CreditcardModel    *models.CreditCard
	CategoryResponse   CategoryResponse
	CreditcardResponse *CreditCardResponse
}

type TransactionResponse struct {
	ID                     uuid.UUID                            `json:"id"`
	Name                   string                               `json:"name"`
	Date                   time.Time                            `json:"date"`
	Value                  float64                              `json:"value"`
	Paid                   bool                                 `json:"paid"`
	Category               CategoryResponse                     `json:"category"`
	Creditcard             *CreditCardResponse                  `json:"creditcard,omitempty"`
	MonthlyTransaction     *ShortMonthlyTransactionResponse     `json:"monthly_transaction,omitempty"`
	AnnualTransaction      *ShortAnnualTransactionResponse      `json:"annual_transaction,omitempty"`
	InstallmentTransaction *ShortInstallmentTransactionResponse `json:"installment_transaction,omitempty"`
	CreatedAt              time.Time                            `json:"created_at"`
	UpdatedAt              time.Time                            `json:"updated_at"`
}

type TransactionRequestFromRecurrentTransaction struct {
	ID uuid.UUID `json:"id"`
}

func (t TransactionRequestFromRecurrentTransaction) Validate() []errors.ApiErrorItem {
	errs := make([]errors.ApiErrorItem, 0)

	if t.ID == uuid.Nil {
		errs = append(errs, errors.InvalidFieldError(constants.InvalidID))
	}

	return errs
}

type TransactionRequest struct {
	Name                     string     `json:"name"`
	Date                     time.Time  `json:"date"`
	Value                    float64    `json:"value"`
	Paid                     bool       `json:"paid"`
	CategoryID               uuid.UUID  `json:"category_id"`
	CreditcardID             *uuid.UUID `json:"creditcard_id,omitempty"`
	MonthlyTransactionID     *uuid.UUID `json:"monthly_transaction_id,omitempty"`
	AnnualTransactionID      *uuid.UUID `json:"annual_transaction_id,omitempty"`
	InstallmentTransactionID *uuid.UUID `json:"installment_transaction_id,omitempty"`
}

func (t TransactionRequest) Validate() []errors.ApiErrorItem {
	errs := make([]errors.ApiErrorItem, 0)

	if utils.IsBlank(t.Name) {
		errs = append(errs, errors.InvalidFieldError(constants.NameEmptyMsg))
	}

	if len(t.Name) >= 255 || len(t.Name) <= 2 {
		errs = append(errs, errors.InvalidFieldError(constants.NameInvalidCharsCountMsg))
	}

	if t.Value < 0 || t.Value >= 1000000000000000.00 {
		errs = append(errs, errors.InvalidFieldError(constants.ValueInvalidMsg))
	}

	if t.Date.IsZero() {
		errs = append(errs, errors.InvalidFieldError(constants.DateEmptyMsg))
	}

	return errs
}
