package dtos

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestMonthlyReportResponseJSONOmitemptyFields(t *testing.T) {
	response := MonthlyReportResponse{
		TotalIncome: 1000,
		TotalDebit:  200,
		TotalCredit: 100,
		Balance:     700,
		Categories:  []CategoriesSpendingResponse{},
		CreditCards: []CreditCardsSpendingResponse{},
	}

	b, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Expected no marshal error, got %v", err)
	}

	jsonString := string(b)
	if strings.Contains(jsonString, "most_spent_category") {
		t.Errorf("Expected most_spent_category to be omitted, got %s", jsonString)
	}
	if strings.Contains(jsonString, "most_spent_creditcard") {
		t.Errorf("Expected most_spent_creditcard to be omitted, got %s", jsonString)
	}
}

func TestMonthlyReportResponseJSONWithMostSpentFields(t *testing.T) {
	category := CategoriesSpendingResponse{ID: uuid.New(), Name: "Food", Value: 100}
	card := CreditCardsSpendingResponse{ID: uuid.New(), Name: "Card", TotalSpent: 300}

	response := MonthlyReportResponse{
		TotalIncome:         1000,
		TotalDebit:          200,
		TotalCredit:         100,
		Balance:             700,
		MostSpentCategory:   &category,
		MostSpentCreditCard: &card,
		Categories:          []CategoriesSpendingResponse{category},
		CreditCards:         []CreditCardsSpendingResponse{card},
	}

	b, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Expected no marshal error, got %v", err)
	}

	jsonString := string(b)
	if !strings.Contains(jsonString, "most_spent_category") {
		t.Errorf("Expected most_spent_category in json, got %s", jsonString)
	}
	if !strings.Contains(jsonString, "most_spent_creditcard") {
		t.Errorf("Expected most_spent_creditcard in json, got %s", jsonString)
	}
}
