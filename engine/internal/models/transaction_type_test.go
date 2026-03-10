package models

import "testing"

func TestTransactionTypeIsValid(t *testing.T) {
	tests := []struct {
		name     string
		tt       TransactionType
		expected bool
	}{
		{"Income is valid", Income, true},
		{"Debit is valid", Debit, true},
		{"Credit is valid", Credit, true},
		{"Invalid transaction type (negative)", TransactionType(-1), false},
		{"Invalid transaction type (too high)", TransactionType(100), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.tt.IsValid()
			if result != tc.expected {
				t.Errorf("Expected IsValid() to return %v for %v, got %v", tc.expected, tc.tt, result)
			}
		})
	}
}

func TestTransactionTypeConstants(t *testing.T) {
	if Income != 0 {
		t.Errorf("Expected Income to be 0, got %d", Income)
	}
	if Debit != 1 {
		t.Errorf("Expected Debit to be 1, got %d", Debit)
	}
	if Credit != 2 {
		t.Errorf("Expected Credit to be 2, got %d", Credit)
	}
}
