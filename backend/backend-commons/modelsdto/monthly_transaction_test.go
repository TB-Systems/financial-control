package modelsdto

import (
	"testing"
	"time"

	"backend-commons/dtos"
	"backend-commons/models"

	"github.com/google/uuid"
)

func TestCreateMonthlyTransactionFromRequest(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditCardID := uuid.New()

	request := dtos.MonthlyTransactionRequest{
		Name:         "Monthly Transaction",
		Value:        500.00,
		Day:          15,
		CategoryID:   categoryID,
		CreditCardID: &creditCardID,
	}

	result := CreateMonthlyTransactionFromRequest(request, userID)

	if result.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, result.UserID)
	}
	if result.Name != request.Name {
		t.Errorf("Expected Name %s, got %s", request.Name, result.Name)
	}
	if result.Value != request.Value {
		t.Errorf("Expected Value %f, got %f", request.Value, result.Value)
	}
	if result.Day != request.Day {
		t.Errorf("Expected Day %d, got %d", request.Day, result.Day)
	}
	if result.CategoryID != request.CategoryID {
		t.Errorf("Expected CategoryID %v, got %v", request.CategoryID, result.CategoryID)
	}
	if *result.CreditCardID != *request.CreditCardID {
		t.Errorf("Expected CreditCardID %v, got %v", *request.CreditCardID, *result.CreditCardID)
	}
}

func TestCreateMonthlyTransactionFromRequestWithoutCreditCard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.MonthlyTransactionRequest{
		Name:         "Monthly Transaction",
		Value:        500.00,
		Day:          15,
		CategoryID:   categoryID,
		CreditCardID: nil,
	}

	result := CreateMonthlyTransactionFromRequest(request, userID)

	if result.CreditCardID != nil {
		t.Errorf("Expected CreditCardID to be nil, got %v", result.CreditCardID)
	}
}

func TestMonthlyTransactionResponseFromModel(t *testing.T) {
	now := time.Now()
	model := models.MonthlyTransaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Monthly Model",
		Value:  750.00,
		Day:    20,
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

	result := MonthlyTransactionResponseFromModel(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.Day != model.Day {
		t.Errorf("Expected Day %d, got %d", model.Day, result.Day)
	}
	if result.Creditcard == nil {
		t.Error("Expected Creditcard to not be nil")
	} else if result.Creditcard.ID != model.Creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", model.Creditcard.ID, result.Creditcard.ID)
	}
}

func TestMonthlyTransactionResponseFromModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.MonthlyTransaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Monthly Without Card",
		Value:  300.00,
		Day:    5,
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

	result := MonthlyTransactionResponseFromModel(model)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestMonthlyTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	model := models.ShortMonthlyTransaction{
		ID:        uuid.New(),
		Name:      "Short Monthly",
		Day:       10,
		Value:     400.00,
		CreatedAt: now,
		UpdatedAt: now,
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

	result := MonthlyTransactionResponseFromShortModel(model, category, creditcard)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.Creditcard.ID != creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", creditcard.ID, result.Creditcard.ID)
	}
}

func TestMonthlyTransactionResponseFromShortModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.ShortMonthlyTransaction{
		ID:        uuid.New(),
		Name:      "Short Monthly No Card",
		Day:       25,
		Value:     600.00,
		CreatedAt: now,
		UpdatedAt: now,
	}

	category := dtos.CategoryResponse{
		ID:              uuid.New(),
		TransactionType: models.Income,
		Name:            "Income",
		Icon:            "income.png",
	}

	result := MonthlyTransactionResponseFromShortModel(model, category, nil)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestShortMonthlyTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	model := models.ShortMonthlyTransaction{
		ID:        uuid.New(),
		Name:      "Short Monthly Only",
		Day:       9,
		Value:     321.65,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ShortMonthlyTransactionResponseFromShortModel(model)

	if result.ID != model.ID || result.Name != model.Name || result.Day != model.Day || result.Value != model.Value {
		t.Errorf("unexpected short monthly response mapping")
	}
}
