package modelsdto

import (
	"testing"
	"time"

	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func TestCreateAnnualTransactionFromRequest(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditCardID := uuid.New()

	request := dtos.AnnualTransactionRequest{
		Name:         "Annual Transaction",
		Value:        1000.00,
		Day:          15,
		Month:        6,
		CategoryID:   categoryID,
		CreditCardID: &creditCardID,
	}

	result := CreateAnnualTransactionFromRequest(request, userID)

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
	if result.Month != request.Month {
		t.Errorf("Expected Month %d, got %d", request.Month, result.Month)
	}
	if result.CategoryID != request.CategoryID {
		t.Errorf("Expected CategoryID %v, got %v", request.CategoryID, result.CategoryID)
	}
	if *result.CreditCardID != *request.CreditCardID {
		t.Errorf("Expected CreditCardID %v, got %v", *request.CreditCardID, *result.CreditCardID)
	}
}

func TestCreateAnnualTransactionFromRequestWithoutCreditCard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.AnnualTransactionRequest{
		Name:         "Annual Transaction",
		Value:        1000.00,
		Day:          15,
		Month:        6,
		CategoryID:   categoryID,
		CreditCardID: nil,
	}

	result := CreateAnnualTransactionFromRequest(request, userID)

	if result.CreditCardID != nil {
		t.Errorf("Expected CreditCardID to be nil, got %v", result.CreditCardID)
	}
}

func TestAnnualTransactionResponseFromModel(t *testing.T) {
	now := time.Now()
	model := models.AnnualTransaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Annual Model",
		Value:  2000.00,
		Day:    20,
		Month:  12,
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

	result := AnnualTransactionResponseFromModel(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.Day != model.Day {
		t.Errorf("Expected Day %d, got %d", model.Day, result.Day)
	}
	if result.Month != model.Month {
		t.Errorf("Expected Month %d, got %d", model.Month, result.Month)
	}
	if result.Creditcard == nil {
		t.Error("Expected Creditcard to not be nil")
	} else if result.Creditcard.ID != model.Creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", model.Creditcard.ID, result.Creditcard.ID)
	}
}

func TestAnnualTransactionResponseFromModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.AnnualTransaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Annual Without Card",
		Value:  1500.00,
		Day:    5,
		Month:  3,
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

	result := AnnualTransactionResponseFromModel(model)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestAnnualTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	model := models.ShortAnnualTransaction{
		ID:        uuid.New(),
		Name:      "Short Annual",
		Day:       10,
		Month:     7,
		Value:     3000.00,
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

	result := AnnualTransactionResponseFromShortModel(model, category, creditcard)

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

func TestAnnualTransactionResponseFromShortModelWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.ShortAnnualTransaction{
		ID:        uuid.New(),
		Name:      "Short Annual No Card",
		Day:       25,
		Month:     11,
		Value:     4000.00,
		CreatedAt: now,
		UpdatedAt: now,
	}

	category := dtos.CategoryResponse{
		ID:              uuid.New(),
		TransactionType: models.Income,
		Name:            "Income",
		Icon:            "income.png",
	}

	result := AnnualTransactionResponseFromShortModel(model, category, nil)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestShortAnnualTransactionResponseFromShortModel(t *testing.T) {
	now := time.Now()
	model := models.ShortAnnualTransaction{
		ID:        uuid.New(),
		Name:      "Short Annual Only",
		Day:       8,
		Month:     2,
		Value:     123.45,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ShortAnnualTransactionResponseFromShortModel(model)

	if result.ID != model.ID || result.Name != model.Name || result.Day != model.Day || result.Month != model.Month || result.Value != model.Value {
		t.Errorf("unexpected short annual response mapping")
	}
}
