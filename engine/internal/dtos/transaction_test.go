package dtos

import (
	"financialcontrol/internal/constants"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTransactionRequestFromRecurrentTransactionValidate_ValidID(t *testing.T) {
	req := TransactionRequestFromRecurrentTransaction{
		ID: uuid.New(),
	}

	err := req.Validate()

	if len(err) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(err), err)
	}
}

func TestTransactionRequestFromRecurrentTransactionValidate_InvalidID(t *testing.T) {
	req := TransactionRequestFromRecurrentTransaction{
		ID: uuid.Nil,
	}

	err := req.Validate()

	if len(err) == 0 {
		t.Fatal("Expected validation error for invalid ID")
	}

	found := false
	for _, validationErr := range err {
		if validationErr.UserMessage == constants.InvalidID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected error with message %q", constants.InvalidID)
	}
}

func TestTransactionRequestValidate_ValidRequest(t *testing.T) {
	req := TransactionRequest{
		Name:       "Grocery Shopping",
		Date:       time.Now(),
		Value:      150.50,
		Paid:       true,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestTransactionRequestValidate_EmptyName(t *testing.T) {
	req := TransactionRequest{
		Name:       "",
		Date:       time.Now(),
		Value:      150.50,
		Paid:       true,
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

func TestTransactionRequestValidate_NameTooShort(t *testing.T) {
	req := TransactionRequest{
		Name:       "AB",
		Date:       time.Now(),
		Value:      150.50,
		Paid:       true,
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

func TestTransactionRequestValidate_NameTooLong(t *testing.T) {
	longName := make([]byte, 255)
	for i := range longName {
		longName[i] = 'a'
	}
	req := TransactionRequest{
		Name:       string(longName),
		Date:       time.Now(),
		Value:      150.50,
		Paid:       true,
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

func TestTransactionRequestValidate_NegativeValue(t *testing.T) {
	req := TransactionRequest{
		Name:       "Shopping",
		Date:       time.Now(),
		Value:      -100,
		Paid:       true,
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

func TestTransactionRequestValidate_ValueTooHigh(t *testing.T) {
	req := TransactionRequest{
		Name:       "Shopping",
		Date:       time.Now(),
		Value:      1000000000000000.00,
		Paid:       true,
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

func TestTransactionRequestValidate_ZeroDate(t *testing.T) {
	req := TransactionRequest{
		Name:       "Shopping",
		Date:       time.Time{},
		Value:      150.50,
		Paid:       true,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.DateEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty date")
	}
}

func TestTransactionRequestValidate_ValidValueEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{"zero value", 0},
		{"small value", 0.01},
		{"large valid value", 999999999999999.00},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := TransactionRequest{
				Name:       "Shopping",
				Date:       time.Now(),
				Value:      tc.value,
				Paid:       true,
				CategoryID: uuid.New(),
			}

			errs := req.Validate()
			for _, err := range errs {
				if err.UserMessage == constants.ValueInvalidMsg {
					t.Errorf("Unexpected error for value %f", tc.value)
				}
			}
		})
	}
}

func TestTransactionRequestValidate_WithOptionalFields(t *testing.T) {
	creditcardID := uuid.New()
	monthlyID := uuid.New()
	annualID := uuid.New()
	installmentID := uuid.New()

	req := TransactionRequest{
		Name:                     "Shopping",
		Date:                     time.Now(),
		Value:                    150.50,
		Paid:                     false,
		CategoryID:               uuid.New(),
		CreditcardID:             &creditcardID,
		MonthlyTransactionID:     &monthlyID,
		AnnualTransactionID:      &annualID,
		InstallmentTransactionID: &installmentID,
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestTransactionRequestValidate_MultipleErrors(t *testing.T) {
	req := TransactionRequest{
		Name:       "",
		Date:       time.Time{},
		Value:      -100,
		Paid:       true,
		CategoryID: uuid.New(),
	}

	errs := req.Validate()
	if len(errs) < 3 {
		t.Errorf("Expected at least 3 errors, got %d", len(errs))
	}
}
