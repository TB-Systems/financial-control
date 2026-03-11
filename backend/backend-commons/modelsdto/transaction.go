package modelsdto

import (
	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func TransactionResponseFromShortTransaction(
	model models.ShortTransaction,
	category dtos.CategoryResponse,
	creditcard *dtos.CreditCardResponse,
	monthlyTransaction *dtos.ShortMonthlyTransactionResponse,
	annualTransaction *dtos.ShortAnnualTransactionResponse,
	installmentTransaction *dtos.ShortInstallmentTransactionResponse,
) dtos.TransactionResponse {
	return dtos.TransactionResponse{
		ID:                     model.ID,
		Name:                   model.Name,
		Date:                   model.Date,
		Value:                  model.Value,
		Paid:                   model.Paid,
		Category:               category,
		Creditcard:             creditcard,
		MonthlyTransaction:     monthlyTransaction,
		AnnualTransaction:      annualTransaction,
		InstallmentTransaction: installmentTransaction,
		CreatedAt:              model.CreatedAt,
		UpdatedAt:              model.UpdatedAt,
	}
}

func CreateTransactionFromTransactionRequest(model dtos.TransactionRequest, userID uuid.UUID) models.CreateTransaction {
	return models.CreateTransaction{
		UserID:                   userID,
		Name:                     model.Name,
		Date:                     model.Date,
		Value:                    model.Value,
		Paid:                     model.Paid,
		CategoryID:               model.CategoryID,
		CreditcardID:             model.CreditcardID,
		MonthlyTransactionID:     model.MonthlyTransactionID,
		AnnualTransactionID:      model.AnnualTransactionID,
		InstallmentTransactionID: model.InstallmentTransactionID,
	}
}

func TransactionResponseFromTransaction(model models.Transaction) dtos.TransactionResponse {
	var creditcard *dtos.CreditCardResponse
	if model.Creditcard != nil {
		creditcardModel := CreditCardResponseFromCreditCard(*model.Creditcard)
		creditcard = &creditcardModel
	}

	var monthlyTransaction *dtos.ShortMonthlyTransactionResponse
	if model.MonthlyTransaction != nil {
		monthlyTransactionResp := ShortMonthlyTransactionResponseFromShortModel(*model.MonthlyTransaction)
		monthlyTransaction = &monthlyTransactionResp
	}

	var annualTransaction *dtos.ShortAnnualTransactionResponse
	if model.AnnualTransaction != nil {
		annualTransactionResp := ShortAnnualTransactionResponseFromShortModel(*model.AnnualTransaction)
		annualTransaction = &annualTransactionResp
	}

	var installmentTransaction *dtos.ShortInstallmentTransactionResponse
	if model.InstallmentTransaction != nil {
		installmentTransactionResp := ShortInstallmentTransactionResponseFromShortModel(*model.InstallmentTransaction)
		installmentTransaction = &installmentTransactionResp
	}

	return dtos.TransactionResponse{
		ID:                     model.ID,
		Name:                   model.Name,
		Date:                   model.Date,
		Value:                  model.Value,
		Paid:                   model.Paid,
		Category:               CategoryResponseFromModel(model.Category),
		Creditcard:             creditcard,
		MonthlyTransaction:     monthlyTransaction,
		AnnualTransaction:      annualTransaction,
		InstallmentTransaction: installmentTransaction,
		CreatedAt:              model.CreatedAt,
		UpdatedAt:              model.UpdatedAt,
	}
}
