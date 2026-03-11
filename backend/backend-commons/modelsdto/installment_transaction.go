package modelsdto

import (
	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func CreateInstallmentTransactionFromRequest(model dtos.InstallmentTransactionRequest, userID uuid.UUID) models.CreateInstallmentTransaction {
	return models.CreateInstallmentTransaction{
		UserID:       userID,
		Name:         model.Name,
		Value:        model.Value,
		InitialDate:  model.InitialDate,
		FinalDate:    model.FinalDate,
		CategoryID:   model.CategoryID,
		CreditCardID: model.CreditCardID,
	}
}

func InstallmentTransactionResponseFromModel(model models.InstallmentTransaction) dtos.InstallmentTransactionResponse {
	var creditcard *dtos.CreditCardResponse
	if model.Creditcard != nil {
		creditcardModel := CreditCardResponseFromCreditCard(*model.Creditcard)
		creditcard = &creditcardModel
	}

	return dtos.InstallmentTransactionResponse{
		ID:          model.ID,
		Name:        model.Name,
		Value:       model.Value,
		InitialDate: model.InitialDate,
		FinalDate:   model.FinalDate,
		Category:    CategoryResponseFromModel(model.Category),
		Creditcard:  creditcard,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func InstallmentTransactionResponseFromShortModel(model models.ShortInstallmentTransaction, category dtos.CategoryResponse, creditcard *dtos.CreditCardResponse) dtos.InstallmentTransactionResponse {
	return dtos.InstallmentTransactionResponse{
		ID:          model.ID,
		Name:        model.Name,
		Value:       model.Value,
		InitialDate: model.InitialDate,
		FinalDate:   model.FinalDate,
		Category:    category,
		Creditcard:  creditcard,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func ShortInstallmentTransactionResponseFromShortModel(model models.ShortInstallmentTransaction) dtos.ShortInstallmentTransactionResponse {
	return dtos.ShortInstallmentTransactionResponse{
		ID:          model.ID,
		Name:        model.Name,
		Value:       model.Value,
		InitialDate: model.InitialDate,
		FinalDate:   model.FinalDate,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
