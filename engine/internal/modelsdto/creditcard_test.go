package modelsdto

import (
	"testing"
	"time"

	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func TestCreditCardResponseFromCreditCard_BasicFields(t *testing.T) {
	model := models.CreditCard{
		ID:               uuid.New(),
		Name:             "Test Card",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         10,
		ExpireDay:        15,
		BackgroundColor:  "#000000",
		TextColor:        "#FFFFFF",
	}

	result := CreditCardResponseFromCreditCard(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.FirstFourNumbers != model.FirstFourNumbers {
		t.Errorf("Expected FirstFourNumbers %s, got %s", model.FirstFourNumbers, result.FirstFourNumbers)
	}
	if result.Limit != model.Limit {
		t.Errorf("Expected Limit %f, got %f", model.Limit, result.Limit)
	}
	if result.CloseDay != model.CloseDay {
		t.Errorf("Expected CloseDay %d, got %d", model.CloseDay, result.CloseDay)
	}
	if result.ExpireDay != model.ExpireDay {
		t.Errorf("Expected ExpireDay %d, got %d", model.ExpireDay, result.ExpireDay)
	}
	if result.BackgroundColor != model.BackgroundColor {
		t.Errorf("Expected BackgroundColor %s, got %s", model.BackgroundColor, result.BackgroundColor)
	}
	if result.TextColor != model.TextColor {
		t.Errorf("Expected TextColor %s, got %s", model.TextColor, result.TextColor)
	}
}

func TestCreateCreditCardFromCreditCardRequest(t *testing.T) {
	userID := uuid.New()
	request := dtos.CreditCardRequest{
		Name:             "New Card",
		FirstFourNumbers: "5678",
		Limit:            10000.00,
		CloseDay:         20,
		ExpireDay:        25,
		BackgroundColor:  "#FF0000",
		TextColor:        "#000000",
	}

	result := CreateCreditCardFromCreditCardRequest(request, userID)

	if result.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, result.UserID)
	}
	if result.Name != request.Name {
		t.Errorf("Expected Name %s, got %s", request.Name, result.Name)
	}
	if result.FirstFourNumbers != request.FirstFourNumbers {
		t.Errorf("Expected FirstFourNumbers %s, got %s", request.FirstFourNumbers, result.FirstFourNumbers)
	}
	if result.Limit != request.Limit {
		t.Errorf("Expected Limit %f, got %f", request.Limit, result.Limit)
	}
	if result.CloseDay != request.CloseDay {
		t.Errorf("Expected CloseDay %d, got %d", request.CloseDay, result.CloseDay)
	}
	if result.ExpireDay != request.ExpireDay {
		t.Errorf("Expected ExpireDay %d, got %d", request.ExpireDay, result.ExpireDay)
	}
	if result.BackgroundColor != request.BackgroundColor {
		t.Errorf("Expected BackgroundColor %s, got %s", request.BackgroundColor, result.BackgroundColor)
	}
	if result.TextColor != request.TextColor {
		t.Errorf("Expected TextColor %s, got %s", request.TextColor, result.TextColor)
	}
}

func TestCreditCardResponseFromCreditCard_WithTimestamps(t *testing.T) {
	now := time.Now()
	model := models.CreditCard{
		ID:               uuid.New(),
		UserID:           uuid.New(),
		Name:             "My Card",
		FirstFourNumbers: "9999",
		Limit:            15000.00,
		CloseDay:         5,
		ExpireDay:        10,
		BackgroundColor:  "#00FF00",
		TextColor:        "#FFFFFF",
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	result := CreditCardResponseFromCreditCard(model)

	if result.ID != model.ID {
		t.Errorf("Expected ID %v, got %v", model.ID, result.ID)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %s, got %s", model.Name, result.Name)
	}
	if result.FirstFourNumbers != model.FirstFourNumbers {
		t.Errorf("Expected FirstFourNumbers %s, got %s", model.FirstFourNumbers, result.FirstFourNumbers)
	}
	if result.Limit != model.Limit {
		t.Errorf("Expected Limit %f, got %f", model.Limit, result.Limit)
	}
	if result.CloseDay != model.CloseDay {
		t.Errorf("Expected CloseDay %d, got %d", model.CloseDay, result.CloseDay)
	}
	if result.ExpireDay != model.ExpireDay {
		t.Errorf("Expected ExpireDay %d, got %d", model.ExpireDay, result.ExpireDay)
	}
	if result.BackgroundColor != model.BackgroundColor {
		t.Errorf("Expected BackgroundColor %s, got %s", model.BackgroundColor, result.BackgroundColor)
	}
	if result.TextColor != model.TextColor {
		t.Errorf("Expected TextColor %s, got %s", model.TextColor, result.TextColor)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
}
