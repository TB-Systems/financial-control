package dtos

import (
	"financialcontrol/internal/constants"
	"testing"
)

func TestCreditCardRequestValidate_ValidRequest(t *testing.T) {
	req := CreditCardRequest{
		Name:             "My Card",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "#FFFFFF",
		TextColor:        "#000000",
	}

	errs := req.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d: %+v", len(errs), errs)
	}
}

func TestCreditCardRequestValidate_EmptyName(t *testing.T) {
	req := CreditCardRequest{
		Name:             "",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "#FFFFFF",
		TextColor:        "#000000",
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

func TestCreditCardRequestValidate_NameTooShort(t *testing.T) {
	req := CreditCardRequest{
		Name:             "AB",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "#FFFFFF",
		TextColor:        "#000000",
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

func TestCreditCardRequestValidate_NameTooLong(t *testing.T) {
	longName := make([]byte, 100)
	for i := range longName {
		longName[i] = 'a'
	}
	req := CreditCardRequest{
		Name:             string(longName),
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "#FFFFFF",
		TextColor:        "#000000",
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

func TestCreditCardRequestValidate_InvalidFirstFourNumbers(t *testing.T) {
	tests := []struct {
		name             string
		firstFourNumbers string
	}{
		{"too short", "123"},
		{"too long", "12345"},
		{"empty", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: tc.firstFourNumbers,
				Limit:            5000.00,
				CloseDay:         15,
				ExpireDay:        25,
				BackgroundColor:  "#FFFFFF",
				TextColor:        "#000000",
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.FirstFourNumbersInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid first four numbers")
			}
		})
	}
}

func TestCreditCardRequestValidate_InvalidLimit(t *testing.T) {
	tests := []struct {
		name  string
		limit float64
	}{
		{"zero", 0},
		{"negative", -100},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            tc.limit,
				CloseDay:         15,
				ExpireDay:        25,
				BackgroundColor:  "#FFFFFF",
				TextColor:        "#000000",
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.LimitInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid limit")
			}
		})
	}
}

func TestCreditCardRequestValidate_InvalidCloseDay(t *testing.T) {
	tests := []struct {
		name     string
		closeDay int32
	}{
		{"zero", 0},
		{"negative", -1},
		{"too high", 32},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            5000.00,
				CloseDay:         tc.closeDay,
				ExpireDay:        25,
				BackgroundColor:  "#FFFFFF",
				TextColor:        "#000000",
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.ClosingDayInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid close day")
			}
		})
	}
}

func TestCreditCardRequestValidate_InvalidExpireDay(t *testing.T) {
	tests := []struct {
		name      string
		expireDay int32
	}{
		{"zero", 0},
		{"negative", -1},
		{"too high", 32},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            5000.00,
				CloseDay:         15,
				ExpireDay:        tc.expireDay,
				BackgroundColor:  "#FFFFFF",
				TextColor:        "#000000",
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.ExpireDayInvalidMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid expire day")
			}
		})
	}
}

func TestCreditCardRequestValidate_EmptyBackgroundColor(t *testing.T) {
	req := CreditCardRequest{
		Name:             "My Card",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "",
		TextColor:        "#000000",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.BackgroundColorEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty background color")
	}
}

func TestCreditCardRequestValidate_InvalidBackgroundColorLength(t *testing.T) {
	tests := []struct {
		name  string
		color string
	}{
		{"too short", "#FFF"},
		{"too long", "#FFFFFFFFFF"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            5000.00,
				CloseDay:         15,
				ExpireDay:        25,
				BackgroundColor:  tc.color,
				TextColor:        "#000000",
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.BackgroundColorInvalidCharsCountMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid background color length")
			}
		})
	}
}

func TestCreditCardRequestValidate_EmptyTextColor(t *testing.T) {
	req := CreditCardRequest{
		Name:             "My Card",
		FirstFourNumbers: "1234",
		Limit:            5000.00,
		CloseDay:         15,
		ExpireDay:        25,
		BackgroundColor:  "#FFFFFF",
		TextColor:        "",
	}

	errs := req.Validate()
	found := false
	for _, err := range errs {
		if err.UserMessage == constants.TextColorEmptyMsg {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected error for empty text color")
	}
}

func TestCreditCardRequestValidate_InvalidTextColorLength(t *testing.T) {
	tests := []struct {
		name  string
		color string
	}{
		{"too short", "#000"},
		{"too long", "#0000000000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            5000.00,
				CloseDay:         15,
				ExpireDay:        25,
				BackgroundColor:  "#FFFFFF",
				TextColor:        tc.color,
			}

			errs := req.Validate()
			found := false
			for _, err := range errs {
				if err.UserMessage == constants.TextColorInvalidCharsCountMsg {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected error for invalid text color length")
			}
		})
	}
}

func TestCreditCardRequestValidate_ValidColorFormats(t *testing.T) {
	tests := []struct {
		name            string
		backgroundColor string
		textColor       string
	}{
		{"7 char hex", "#FFFFFF", "#000000"},
		{"9 char hex with alpha", "#FFFFFFFF", "#000000FF"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := CreditCardRequest{
				Name:             "My Card",
				FirstFourNumbers: "1234",
				Limit:            5000.00,
				CloseDay:         15,
				ExpireDay:        25,
				BackgroundColor:  tc.backgroundColor,
				TextColor:        tc.textColor,
			}

			errs := req.Validate()
			// Filter out color errors
			colorErrors := 0
			for _, err := range errs {
				if err.UserMessage == constants.BackgroundColorInvalidCharsCountMsg ||
					err.UserMessage == constants.TextColorInvalidCharsCountMsg {
					colorErrors++
				}
			}
			if colorErrors > 0 {
				t.Errorf("Expected no color errors, got %d", colorErrors)
			}
		})
	}
}
