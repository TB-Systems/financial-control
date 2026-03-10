package modelsdto

import (
	"testing"
	"time"

	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func TestTransactionResponseFromShortTransaction(t *testing.T) {
	now := time.Now()
	model := models.ShortTransaction{
		ID:        uuid.New(),
		Name:      "Test Transaction",
		Date:      now,
		Value:     100.00,
		Paid:      true,
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
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         10,
		ExpireDay:        15,
		BackgroundColor:  "#000000",
		TextColor:        "#FFFFFF",
	}

	result := TransactionResponseFromShortTransaction(model, category, creditcard, nil, nil, nil)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Date != model.Date {
		t.Errorf("Expected Date %v, got %v", model.Date, result.Date)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %f, got %f", model.Value, result.Value)
	}
	if result.Paid != model.Paid {
		t.Errorf("Expected Paid %v, got %v", model.Paid, result.Paid)
	}
	if result.Category.ID != category.ID {
		t.Errorf("Expected Category.ID %v, got %v", category.ID, result.Category.ID)
	}
	if result.Creditcard.ID != creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", creditcard.ID, result.Creditcard.ID)
	}
}

func TestTransactionResponseFromShortTransactionWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.ShortTransaction{
		ID:        uuid.New(),
		Name:      "Test Transaction",
		Date:      now,
		Value:     100.00,
		Paid:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	category := dtos.CategoryResponse{
		ID:              uuid.New(),
		TransactionType: models.Income,
		Name:            "Category",
		Icon:            "icon.png",
	}

	result := TransactionResponseFromShortTransaction(model, category, nil, nil, nil, nil)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}
}

func TestCreateTransactionFromTransactionRequest(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()
	monthlyID := uuid.New()
	annualID := uuid.New()
	installmentID := uuid.New()
	now := time.Now()

	request := dtos.TransactionRequest{
		Name:                     "New Transaction",
		Date:                     now,
		Value:                    250.00,
		Paid:                     true,
		CategoryID:               categoryID,
		CreditcardID:             &creditcardID,
		MonthlyTransactionID:     &monthlyID,
		AnnualTransactionID:      &annualID,
		InstallmentTransactionID: &installmentID,
	}

	result := CreateTransactionFromTransactionRequest(request, userID)

	if result.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, result.UserID)
	}
	if result.Name != request.Name {
		t.Errorf("Expected Name %s, got %s", request.Name, result.Name)
	}
	if result.Date != request.Date {
		t.Errorf("Expected Date %v, got %v", request.Date, result.Date)
	}
	if result.Value != request.Value {
		t.Errorf("Expected Value %f, got %f", request.Value, result.Value)
	}
	if result.Paid != request.Paid {
		t.Errorf("Expected Paid %v, got %v", request.Paid, result.Paid)
	}
	if result.CategoryID != request.CategoryID {
		t.Errorf("Expected CategoryID %v, got %v", request.CategoryID, result.CategoryID)
	}
	if *result.CreditcardID != *request.CreditcardID {
		t.Errorf("Expected CreditcardID %v, got %v", *request.CreditcardID, *result.CreditcardID)
	}
}

func TestTransactionResponseFromTransaction(t *testing.T) {
	now := time.Now()
	model := models.Transaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Full Transaction",
		Date:   now,
		Value:  500.00,
		Paid:   true,
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
			Limit:            10000.00,
			CloseDay:         10,
			ExpireDay:        15,
			BackgroundColor:  "#000000",
			TextColor:        "#FFFFFF",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := TransactionResponseFromTransaction(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Creditcard == nil {
		t.Error("Expected Creditcard to not be nil")
	} else if result.Creditcard.ID != model.Creditcard.ID {
		t.Errorf("Expected Creditcard.ID %v, got %v", model.Creditcard.ID, result.Creditcard.ID)
	}
}

func TestTransactionResponseFromTransactionWithoutCreditCard(t *testing.T) {
	now := time.Now()
	model := models.Transaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Transaction Without Card",
		Date:   now,
		Value:  300.00,
		Paid:   false,
		Category: models.Category{
			ID:              uuid.New(),
			TransactionType: models.Income,
			Name:            "Income Category",
			Icon:            "income.png",
		},
		Creditcard: nil,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	result := TransactionResponseFromTransaction(model)

	if result.Creditcard != nil {
		t.Errorf("Expected Creditcard to be nil, got %v", result.Creditcard)
	}

	if result.MonthlyTransaction != nil {
		t.Errorf("Expected MonthlyTransaction to be nil, got %v", result.MonthlyTransaction)
	}

	if result.AnnualTransaction != nil {
		t.Errorf("Expected AnnualTransaction to be nil, got %v", result.AnnualTransaction)
	}

	if result.InstallmentTransaction != nil {
		t.Errorf("Expected InstallmentTransaction to be nil, got %v", result.InstallmentTransaction)
	}
}

func TestTransactionResponseFromTransactionWithRecurrentRelations(t *testing.T) {
	now := time.Now()
	monthlyID := uuid.New()
	annualID := uuid.New()
	installmentID := uuid.New()
	creditCardID := uuid.New()

	model := models.Transaction{
		ID:     uuid.New(),
		UserID: uuid.New(),
		Name:   "Transaction With Relations",
		Date:   now,
		Value:  420.00,
		Paid:   true,
		Category: models.Category{
			ID:              uuid.New(),
			TransactionType: models.Debit,
			Name:            "Category",
			Icon:            "icon.png",
		},
		Creditcard: &models.CreditCard{
			ID:               creditCardID,
			Name:             "Card",
			FirstFourNumbers: "9999",
			Limit:            2000,
			CloseDay:         8,
			ExpireDay:        20,
			BackgroundColor:  "#101010",
			TextColor:        "#fafafa",
		},
		MonthlyTransaction: &models.ShortMonthlyTransaction{
			ID:           monthlyID,
			UserID:       uuid.New(),
			Name:         "Monthly",
			Day:          10,
			Value:        100,
			CategoryID:   uuid.New(),
			CreditCardID: &creditCardID,
		},
		AnnualTransaction: &models.ShortAnnualTransaction{
			ID:           annualID,
			UserID:       uuid.New(),
			Name:         "Annual",
			Day:          12,
			Month:        6,
			Value:        1200,
			CategoryID:   uuid.New(),
			CreditCardID: &creditCardID,
		},
		InstallmentTransaction: &models.ShortInstallmentTransaction{
			ID:           installmentID,
			UserID:       uuid.New(),
			Name:         "Installment",
			InitialDate:  now,
			FinalDate:    now.AddDate(0, 3, 0),
			Value:        140,
			CategoryID:   uuid.New(),
			CreditCardID: &creditCardID,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := TransactionResponseFromTransaction(model)

	if result.MonthlyTransaction == nil || result.MonthlyTransaction.ID != monthlyID {
		t.Errorf("Expected MonthlyTransaction ID %v", monthlyID)
	}

	if result.AnnualTransaction == nil || result.AnnualTransaction.ID != annualID {
		t.Errorf("Expected AnnualTransaction ID %v", annualID)
	}

	if result.InstallmentTransaction == nil || result.InstallmentTransaction.ID != installmentID {
		t.Errorf("Expected InstallmentTransaction ID %v", installmentID)
	}
}
