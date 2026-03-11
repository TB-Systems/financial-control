package dtos

import (
	"backend-commons/constants"
	"testing"

	"github.com/google/uuid"
)

func TestMonthlyTransactionRequestValidate_ValidRequest(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "Rent Payment",
		Value:      1500.00,
		Day:        5,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestMonthlyTransactionRequestValidate_EmptyName(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "",
		Value:      1500.00,
		Day:        5,
		CategoryID: uuid.New(),
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

func TestMonthlyTransactionRequestValidate_NameTooShort(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "A",
		Value:      1500.00,
		Day:        5,
		CategoryID: uuid.New(),
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

func TestMonthlyTransactionRequestValidate_NameTooLong(t *testing.T) {
	longName := make([]byte, 256)
	for i := range longName {
		longName[i] = 'a'
	}
	req := MonthlyTransactionRequest{
		Name:       string(longName),
		Value:      1500.00,
		Day:        5,
		CategoryID: uuid.New(),
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

func TestMonthlyTransactionRequestValidate_NegativeValue(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "Rent Payment",
		Value:      -100,
		Day:        5,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.ValueInvalidMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for negative value")
	}
}

func TestMonthlyTransactionRequestValidate_ValueTooHigh(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "Rent Payment",
		Value:      1000000000000000.00,
		Day:        5,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.ValueInvalidMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for value too high")
	}
}

func TestMonthlyTransactionRequestValidate_InvalidDay(t *testing.T) {
	tests := []struct {
		name string
		day  int32
	}{
		{"zero", 0},
		{"negative", -1},
		{"too high", 32},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := MonthlyTransactionRequest{
				Name:       "Rent Payment",
				Value:      1500.00,
				Day:        tc.day,
				CategoryID: uuid.New(),
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.DayInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid day %d", tc.day)
			}
		})
	}
}

func TestMonthlyTransactionRequestValidate_ValidDayEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		day  int32
	}{
		{"first day", 1},
		{"last day", 31},
		{"mid month", 15},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := MonthlyTransactionRequest{
				Name:       "Rent Payment",
				Value:      1500.00,
				Day:        tc.day,
				CategoryID: uuid.New(),
			}

			errs := req.Validate()
			for _, err := range errs {
				if err.UserMessage == constants.DayInvalidMsg {
					t.Errorf("Unexpected error for valid day %d", tc.day)
				}
			}
		})
	}
}

func TestMonthlyTransactionRequestValidate_WithCreditCard(t *testing.T) {
	creditCardID := uuid.New()
	req := MonthlyTransactionRequest{
		Name:         "Credit Card Bill",
		Value:        500.00,
		Day:          10,
		CategoryID:   uuid.New(),
		CreditCardID: &creditCardID,
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestMonthlyTransactionRequestValidate_MultipleErrors(t *testing.T) {
	req := MonthlyTransactionRequest{
		Name:       "",
		Value:      -100,
		Day:        0,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) < 3 {
		t.Errorf("Expected at least 3 errors, got %d", len(errs))
	}
}
