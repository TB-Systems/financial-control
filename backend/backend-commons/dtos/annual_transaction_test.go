package dtos

import (
	"backend-commons/constants"
	"testing"

	"github.com/google/uuid"
)

func TestAnnualTransactionRequestValidate_ValidRequest(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "Annual Insurance",
		Value:      2000.00,
		Day:        15,
		Month:      6,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestAnnualTransactionRequestValidate_EmptyName(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "",
		Value:      2000.00,
		Day:        15,
		Month:      6,
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

func TestAnnualTransactionRequestValidate_NameTooShort(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "A",
		Value:      2000.00,
		Day:        15,
		Month:      6,
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

func TestAnnualTransactionRequestValidate_NameTooLong(t *testing.T) {
	longName := make([]byte, 256)
	for i := range longName {
		longName[i] = 'a'
	}
	req := AnnualTransactionRequest{
		Name:       string(longName),
		Value:      2000.00,
		Day:        15,
		Month:      6,
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

func TestAnnualTransactionRequestValidate_NegativeValue(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "Annual Insurance",
		Value:      -100,
		Day:        15,
		Month:      6,
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

func TestAnnualTransactionRequestValidate_ValueTooHigh(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "Annual Insurance",
		Value:      1000000000000000.00,
		Day:        15,
		Month:      6,
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

func TestAnnualTransactionRequestValidate_InvalidDay(t *testing.T) {
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
			req := AnnualTransactionRequest{
				Name:       "Annual Insurance",
				Value:      2000.00,
				Day:        tc.day,
				Month:      6,
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

func TestAnnualTransactionRequestValidate_InvalidMonth(t *testing.T) {
	tests := []struct {
		name  string
		month int32
	}{
		{"zero", 0},
		{"negative", -1},
		{"too high", 13},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := AnnualTransactionRequest{
				Name:       "Annual Insurance",
				Value:      2000.00,
				Day:        15,
				Month:      tc.month,
				CategoryID: uuid.New(),
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.MonthInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid month %d", tc.month)
			}
		})
	}
}

func TestAnnualTransactionRequestValidate_ValidMonthEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		month int32
	}{
		{"January", 1},
		{"December", 12},
		{"mid year", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := AnnualTransactionRequest{
				Name:       "Annual Insurance",
				Value:      2000.00,
				Day:        15,
				Month:      tc.month,
				CategoryID: uuid.New(),
			}

			errs := req.Validate()
			for _, err := range errs {
				if err.UserMessage == constants.MonthInvalidMsg {
					t.Errorf("Unexpected error for valid month %d", tc.month)
				}
			}
		})
	}
}

func TestAnnualTransactionRequestValidate_WithCreditCard(t *testing.T) {
	creditCardID := uuid.New()
	req := AnnualTransactionRequest{
		Name:         "Annual Insurance",
		Value:        2000.00,
		Day:          15,
		Month:        6,
		CategoryID:   uuid.New(),
		CreditCardID: &creditCardID,
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestAnnualTransactionRequestValidate_MultipleErrors(t *testing.T) {
	req := AnnualTransactionRequest{
		Name:       "",
		Value:      -100,
		Day:        0,
		Month:      0,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) < 4 {
		t.Errorf("Expected at least 4 errors, got %d", len(errs))
	}
}
