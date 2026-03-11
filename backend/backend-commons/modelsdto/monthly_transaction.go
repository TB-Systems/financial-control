package modelsdto

import (
	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func CreateMonthlyTransactionFromRequest(model dtos.MonthlyTransactionRequest, userID uuid.UUID) models.CreateMonthlyTransaction {
	return models.CreateMonthlyTransaction{
		UserID:       userID,
		Name:         model.Name,
		Value:        model.Value,
		Day:          model.Day,
		CategoryID:   model.CategoryID,
		CreditCardID: model.CreditCardID,
	}
}

func MonthlyTransactionResponseFromModel(model models.MonthlyTransaction) dtos.MonthlyTransactionResponse {
	var creditcard *dtos.CreditCardResponse
	if model.Creditcard != nil {
		creditcardModel := CreditCardResponseFromCreditCard(*model.Creditcard)
		creditcard = &creditcardModel
	}

	return dtos.MonthlyTransactionResponse{
		ID:         model.ID,
		Name:       model.Name,
		Value:      model.Value,
		Day:        model.Day,
		Category:   CategoryResponseFromModel(model.Category),
		Creditcard: creditcard,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func MonthlyTransactionResponseFromShortModel(model models.ShortMonthlyTransaction, category dtos.CategoryResponse, creditcard *dtos.CreditCardResponse) dtos.MonthlyTransactionResponse {
	return dtos.MonthlyTransactionResponse{
		ID:         model.ID,
		Name:       model.Name,
		Value:      model.Value,
		Day:        model.Day,
		Category:   category,
		Creditcard: creditcard,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func ShortMonthlyTransactionResponseFromShortModel(model models.ShortMonthlyTransaction) dtos.ShortMonthlyTransactionResponse {
	return dtos.ShortMonthlyTransactionResponse{
		ID:        model.ID,
		Name:      model.Name,
		Value:     model.Value,
		Day:       model.Day,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
