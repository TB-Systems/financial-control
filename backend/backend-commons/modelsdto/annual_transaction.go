package modelsdto

import (
	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func CreateAnnualTransactionFromRequest(model dtos.AnnualTransactionRequest, userID uuid.UUID) models.CreateAnnualTransaction {
	return models.CreateAnnualTransaction{
		UserID:       userID,
		Name:         model.Name,
		Value:        model.Value,
		Day:          model.Day,
		Month:        model.Month,
		CategoryID:   model.CategoryID,
		CreditCardID: model.CreditCardID,
	}
}

func AnnualTransactionResponseFromModel(model models.AnnualTransaction) dtos.AnnualTransactionResponse {
	var creditcard *dtos.CreditCardResponse
	if model.Creditcard != nil {
		creditcardModel := CreditCardResponseFromCreditCard(*model.Creditcard)
		creditcard = &creditcardModel
	}

	return dtos.AnnualTransactionResponse{
		ID:         model.ID,
		Name:       model.Name,
		Value:      model.Value,
		Day:        model.Day,
		Month:      model.Month,
		Category:   CategoryResponseFromModel(model.Category),
		Creditcard: creditcard,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func AnnualTransactionResponseFromShortModel(model models.ShortAnnualTransaction, category dtos.CategoryResponse, creditcard *dtos.CreditCardResponse) dtos.AnnualTransactionResponse {
	return dtos.AnnualTransactionResponse{
		ID:         model.ID,
		Name:       model.Name,
		Value:      model.Value,
		Day:        model.Day,
		Month:      model.Month,
		Category:   category,
		Creditcard: creditcard,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func ShortAnnualTransactionResponseFromShortModel(model models.ShortAnnualTransaction) dtos.ShortAnnualTransactionResponse {
	return dtos.ShortAnnualTransactionResponse{
		ID:        model.ID,
		Name:      model.Name,
		Value:     model.Value,
		Day:       model.Day,
		Month:     model.Month,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
