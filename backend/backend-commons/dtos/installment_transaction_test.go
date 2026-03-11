package dtos

import (
	"backend-commons/constants"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestInstallmentTransactionRequestValidate_ValidRequest(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       500.00,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestInstallmentTransactionRequestValidate_EmptyName(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "",
		Value:       500.00,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
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

func TestInstallmentTransactionRequestValidate_NameTooShort(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "A",
		Value:       500.00,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
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

func TestInstallmentTransactionRequestValidate_NameTooLong(t *testing.T) {
	longName := make([]byte, 256)
	for i := range longName {
		longName[i] = 'a'
	}
	req := InstallmentTransactionRequest{
		Name:        string(longName),
		Value:       500.00,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
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

func TestInstallmentTransactionRequestValidate_NegativeValue(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       -100,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
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

func TestInstallmentTransactionRequestValidate_ValueTooHigh(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       1000000000000000.00,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
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

func TestInstallmentTransactionRequestValidate_EmptyInitialDate(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       500.00,
		InitialDate: time.Time{},
		FinalDate:   time.Now().AddDate(0, 12, 0),
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.InitialDateEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty initial date")
	}
}

func TestInstallmentTransactionRequestValidate_EmptyFinalDate(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       500.00,
		InitialDate: time.Now(),
		FinalDate:   time.Time{},
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.FinalDateEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty final date")
	}
}

func TestInstallmentTransactionRequestValidate_FinalDateBeforeInitialDate(t *testing.T) {
	initialDate := time.Now()
	finalDate := initialDate.AddDate(0, -1, 0) // 1 month before

	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       500.00,
		InitialDate: initialDate,
		FinalDate:   finalDate,
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.FinalDateBeforeInitialDateMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for final date before initial date")
	}
}

func TestInstallmentTransactionRequestValidate_InitialDateEqualsFinalDate(t *testing.T) {
	sameDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)

	req := InstallmentTransactionRequest{
		Name:        "Furniture Purchase",
		Value:       500.00,
		InitialDate: sameDate,
		FinalDate:   sameDate,
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.InitialDateEqualsFinalDateMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for initial date equals final date")
	}
}

func TestInstallmentTransactionRequestValidate_WithCreditCard(t *testing.T) {
	creditCardID := uuid.New()
	req := InstallmentTransactionRequest{
		Name:         "Furniture Purchase",
		Value:        500.00,
		InitialDate:  time.Now(),
		FinalDate:    time.Now().AddDate(0, 12, 0),
		CategoryID:   uuid.New(),
		CreditCardID: &creditCardID,
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestInstallmentTransactionRequestValidate_MultipleErrors(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "",
		Value:       -100,
		InitialDate: time.Time{},
		FinalDate:   time.Time{},
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	if len(errs) < 4 {
		t.Errorf("Expected at least 4 errors, got %d", len(errs))
	}
}

func TestInstallmentTransactionRequestValidate_BothDatesZeroNoDateOrderError(t *testing.T) {
	req := InstallmentTransactionRequest{
		Name:        "Test",
		Value:       100.00,
		InitialDate: time.Time{},
		FinalDate:   time.Time{},
		CategoryID:  uuid.New(),
	}

	errs := req.Validate()
	// Should not have "final date before initial date" error when both dates are zero
	for _, err := range errs {
		if err.UserMessage == constants.FinalDateBeforeInitialDateMsg {
			t.Errorf("Should not have 'final date before initial date' error when both dates are zero")
		}
	}
}
