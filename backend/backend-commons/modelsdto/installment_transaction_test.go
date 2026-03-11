package modelsdto

import (
	"testing"
	"time"

	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func TestCreateInstallmentTransactionFromRequest(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditCardID := uuid.New()
	initialDate := time.Now()
	finalDate := initialDate.AddDate(0, 12, 0)

	request := dtos.InstallmentTransactionRequest{
		Name:         "Installment Transaction",
		Value:        1200.00,
		InitialDate:  initialDate,
		FinalDate:    finalDate,
		CategoryID:   categoryID,
		CreditCardID: &creditCardID,
	}

	result := CreateInstallmentTransactionFromRequest(request, userID)

	if result.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, result.UserID)
	}
	if result.Name != request.Name {
		t.Errorf("Expected Name %s, got %s", request.Name, result.Name)
	}
	if result.Value != request.Value {
		t.Errorf("Expected Value %f, got %f", request.Value, result.Value)
	}
	if result.InitialDate != request.InitialDate {
		t.Errorf("Expected InitialDate %v, got %v", request.InitialDate, result.InitialDate)
	}
	if result.FinalDate != request.FinalDate {
		t.Errorf("Expected FinalDate %v, got %v", request.FinalDate, result.FinalDate)
	}
	if result.CategoryID != request.CategoryID {
		t.Errorf("Expected CategoryID %v, got %v", request.CategoryID, result.CategoryID)
	}
	if *result.CreditCardID != *request.CreditCardID {
		t.Errorf("Expected CreditCardID %v, got %v", *request.CreditCardID, *result.CreditCardID)
	}
}

func TestCreateInstallmentTransactionFromRequestWithoutCreditCard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	initialDate := time.Now()
	finalDate := initialDate.AddDate(0, 6, 0)

	request := dtos.InstallmentTransactionRequest{
		Name:         "Installment Transaction",
		Value:        600.00,
		InitialDate:  initialDate,
		FinalDate:    finalDate,
		CategoryID:   categoryID,
		CreditCardID: nil,
	}

	result := CreateInstallmentTransactionFromRequest(request, userID)

	if result.CreditCardID != nil {
		t.Errorf("Expected CreditCardID to be nil, got %v", result.CreditCardID)
	}
}

func TestInstallmentTransactionResponseFromModel(t *testing.T) {
	now := time.Now()
	finalDate := now.AddDate(0, 12, 0)
	model := models.InstallmentTransaction{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Name:        "Installment Model",
		Value:       2400.00,
		InitialDate: now,
		FinalDate:   finalDate,
		Category: models.Category{
			ID:              uuid.New(),
			TransactionType: models.Debit,
			Name:            "Category",
			Icon:            "icon.png",
		},
		Creditcard: &models.CreditCard{
			ID:               uuid.New(),
			Name:             "Card",
			FirstFourNumbers: "1234",
			Limit:            5000.00,
			CloseDay:         10,
			ExpireDay:        15,
			BackgroundColor:  "#000000",
			TextColor:        "#FFFFFF",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := InstallmentTransactionResponseFromModel(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.InitialDate != model.InitialDate {
		t.Errorf("Expected InitialDate %v, got %v", model.InitialDate, result.InitialDate)
	}
	if result.FinalDate != model.FinalDate {
		t.Errorf("Expected FinalDate %v, got %v", model.FinalDate, result.FinalDate)
	}
	if result.Creditcard == nil {
		t.Error("Expected Creditcard to not be nil")
	} else if result.Creditcard.ID != model.Creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", model.Creditcard.ID, result.Creditcard.ID)
	}
}

func TestInstallmentTransactionResponseFromModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	finalDate := now.AddDate(0, 6, 0)
	model := models.InstallmentTransaction{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Name:        "Installment Without Card",
		Value:       1800.00,
		InitialDate: now,
		FinalDate:   finalDate,
		Category: models.Category{
			ID:              uuid.New(),
			TransactionType: models.Income,
			Name:            "Income",
			Icon:            "income.png",
		},
		Creditcard: nil,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	result := InstallmentTransactionResponseFromModel(model)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestInstallmentTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	finalDate := now.AddDate(0, 10, 0)
	model := models.ShortInstallmentTransaction{
		ID:          uuid.New(),
		Name:        "Short Installment",
		Value:       3000.00,
		InitialDate: now,
		FinalDate:   finalDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	category := dtos.CategoryResponse{
		ID:              uuid.New(),
		TransactionType: models.Debit,
		Name:            "Category",
		Icon:            "icon.png",
	}

	creditcard := &dtos.CreditCardResponse{
		ID:               uuid.New(),
		Name:             "Card",
		FirstFourNumbers: "5678",
		Limit:            8000.00,
		CloseDay:         15,
		ExpireDay:        20,
		BackgroundColor:  "#FF0000",
		TextColor:        "#FFFFFF",
	}

	result := InstallmentTransactionResponseFromShortModel(model, category, creditcard)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.Creditcard.ID != creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", creditcard.ID, result.Creditcard.ID)
	}
}

func TestInstallmentTransactionResponseFromShortModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	finalDate := now.AddDate(0, 8, 0)
	model := models.ShortInstallmentTransaction{
		ID:          uuid.New(),
		Name:        "Short Installment No Card",
		Value:       2500.00,
		InitialDate: now,
		FinalDate:   finalDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	category := dtos.CategoryResponse{
		ID:              uuid.New(),
		TransactionType: models.Income,
		Name:            "Income",
		Icon:            "income.png",
	}

	result := InstallmentTransactionResponseFromShortModel(model, category, nil)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestShortInstallmentTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	finalDate := now.AddDate(0, 4, 0)
	model := models.ShortInstallmentTransaction{
		ID:          uuid.New(),
		Name:        "Short Installment Only",
		Value:       789.10,
		InitialDate: now,
		FinalDate:   finalDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := ShortInstallmentTransactionResponseFromShortModel(model)

	if result.ID != model.ID || result.Name != model.Name || result.Value != model.Value || result.InitialDate != model.InitialDate || result.FinalDate != model.FinalDate {
		t.Errorf("unexpected short installment response mapping")
	}
}
