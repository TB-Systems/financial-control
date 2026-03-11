package modelsdto

import (
	"testing"

	"backend-commons/models"

	"github.com/google/uuid"
)

func TestMonthlyReportResponseFromModelsWithData(t *testing.T) {
	categoryID := uuid.New()
	cardID := uuid.New()

	model := models.MonthlyReport{
		TotalIncome: 1000,
		TotalDebit:  300,
		TotalCredit: 100,
		Balance:     600,
	}

	categories := []models.CategoriesSpending{
		{
			CategoryID:              categoryID,
			CategoryName:            "Food",
			CategoryIcon:            "restaurant",
			CategoryTransactionType: models.Debit,
			TotalSpent:              250,
		},
	}

	cards := []models.CreditCardsSpending{
		{
			ID:               cardID,
			Name:             "Main Card",
			FirstFourNumbers: "1234",
			Limit:            4000,
			CloseDay:         10,
			ExpireDay:        20,
			BackgroundColor:  "#000",
			TextColor:        "#fff",
			TotalSpent:       500,
		},
	}

	response := MonthlyReportResponseFromModels(model, categories, cards)

	if response.TotalIncome != 1000 || response.Balance != 600 {
		t.Errorf("Unexpected summary values: %+v", response)
	}

	if len(response.Categories) != 1 || response.Categories[0].ID != categoryID {
		t.Errorf("Unexpected categories mapping: %+v", response.Categories)
	}

	if len(response.CreditCards) != 1 || response.CreditCards[0].ID != cardID {
		t.Errorf("Unexpected credit cards mapping: %+v", response.CreditCards)
	}

	if response.MostSpentCategory == nil || response.MostSpentCategory.ID != categoryID {
		t.Errorf("Expected MostSpentCategory to be first category")
	}

	if response.MostSpentCreditCard == nil || response.MostSpentCreditCard.ID != cardID {
		t.Errorf("Expected MostSpentCreditCard to be first card")
	}
}

func TestMonthlyReportResponseFromModelsWithoutData(t *testing.T) {
	response := MonthlyReportResponseFromModels(
		models.MonthlyReport{TotalIncome: 10, Balance: 10},
		[]models.CategoriesSpending{},
		[]models.CreditCardsSpending{},
	)

	if response.MostSpentCategory != nil {
		t.Errorf("Expected nil MostSpentCategory")
	}

	if response.MostSpentCreditCard != nil {
		t.Errorf("Expected nil MostSpentCreditCard")
	}

	if len(response.Categories) != 0 || len(response.CreditCards) != 0 {
		t.Errorf("Expected empty slices")
	}
}
