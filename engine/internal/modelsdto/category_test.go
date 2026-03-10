package modelsdto

import (
	"testing"
	"time"

	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func TestCreateCategoryFromRequest(t *testing.T) {
	userID := uuid.New()
	transactionType := models.Debit

	request := dtos.CategoryRequest{
		TransactionType: &transactionType,
		Name:            "Test Category",
		Icon:            "icon.png",
	}

	result := CreateCategoryFromRequest(request, userID)

	if result.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, result.UserID)
	}
	if result.TransactionType != transactionType {
		t.Errorf("Expected TransactionType %v, got %v", transactionType, result.TransactionType)
	}
	if result.Name != request.Name {
		t.Errorf("Expected Name %s, got %s", request.Name, result.Name)
	}
	if result.Icon != request.Icon {
		t.Errorf("Expected Icon %s, got %s", request.Icon, result.Icon)
	}
}

func TestCategoryResponseFromModel_WithTimestamps(t *testing.T) {
	now := time.Now()
	model := models.Category{
		ID:              uuid.New(),
		UserID:          uuid.New(),
		TransactionType: models.Income,
		Name:            "Test Category",
		Icon:            "icon.png",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	result := CategoryResponseFromModel(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.TransactionType != model.TransactionType {
		t.Errorf("Expected TransactionType %v, got %v", model.TransactionType, result.TransactionType)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Icon != model.Icon {
		t.Errorf("Expected Icon %s, got %s", model.Icon, result.Icon)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
}

func TestCategoryResponseFromModel_BasicFields(t *testing.T) {
	model := models.Category{
		ID:              uuid.New(),
		UserID:          uuid.New(),
		TransactionType: models.Debit,
		Name:            "Short Category",
		Icon:            "short-icon.png",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	result := CategoryResponseFromModel(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.TransactionType != model.TransactionType {
		t.Errorf("Expected TransactionType %v, got %v", model.TransactionType, result.TransactionType)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.Icon != model.Icon {
		t.Errorf("Expected Icon %s, got %s", model.Icon, result.Icon)
	}
}
