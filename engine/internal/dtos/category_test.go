package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/models"
	"testing"
)

func TestCategoryRequestValidate_ValidRequest(t *testing.T) {
	tt := models.Income
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "Food",
		Icon:            "🍔",
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestCategoryRequestValidate_NilTransactionType(t *testing.T) {
	req := CategoryRequest{
		TransactionType: nil,
		Name:            "Food",
		Icon:            "🍔",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.TransactionTypeEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for nil transaction type")
	}
}

func TestCategoryRequestValidate_InvalidTransactionType(t *testing.T) {
	tt := models.TransactionType(99)
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "Food",
		Icon:            "🍔",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.TransactionTypeMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for invalid transaction type")
	}
}

func TestCategoryRequestValidate_EmptyName(t *testing.T) {
	tt := models.Income
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "",
		Icon:            "🍔",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.NameEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty name")
	}
}

func TestCategoryRequestValidate_NameTooShort(t *testing.T) {
	tt := models.Income
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "A",
		Icon:            "🍔",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.NameInvalidCharsCountMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for name too short")
	}
}

func TestCategoryRequestValidate_NameTooLong(t *testing.T) {
	tt := models.Income
	longName := make([]byte, 256)
	for i := range longName {
		longName[i] = 'a'
	}
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            string(longName),
		Icon:            "🍔",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.NameInvalidCharsCountMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for name too long")
	}
}

func TestCategoryRequestValidate_EmptyIcon(t *testing.T) {
	tt := models.Income
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "Food",
		Icon:            "",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.IconEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty icon")
	}
}

func TestCategoryRequestValidate_IconTooShort(t *testing.T) {
	tt := models.Income
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "Food",
		Icon:            "A",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.IconInvalidCharsCountMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for icon too short")
	}
}

func TestCategoryRequestValidate_IconTooLong(t *testing.T) {
	tt := models.Income
	longIcon := make([]byte, 256)
	for i := range longIcon {
		longIcon[i] = 'a'
	}
	req := CategoryRequest{
		TransactionType: &tt,
		Name:            "Food",
		Icon:            string(longIcon),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.IconInvalidCharsCountMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for icon too long")
	}
}

func TestCategoryRequestValidate_MultipleErrors(t *testing.T) {
	req := CategoryRequest{
		TransactionType: nil,
		Name:            "",
		Icon:            "",
	}

	errs := req.Validate()
	if len(errs) < 3 {
		t.Errorf("Expected at least 3 errors, got %d", len(errs))
	}
}
