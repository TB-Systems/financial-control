package modelsdto

import (
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"
)

func MonthlyReportResponseFromModels(model models.MonthlyReport, spendingCategories []models.CategoriesSpending, spendingCreditCards []models.CreditCardsSpending) dtos.MonthlyReportResponse {
	categories := make([]dtos.CategoriesSpendingResponse, len(spendingCategories))

	for i, category := range spendingCategories {
		categories[i] = dtos.CategoriesSpendingResponse{
			ID:              category.CategoryID,
			Name:            category.CategoryName,
			Icon:            category.CategoryIcon,
			TransactionType: int(category.CategoryTransactionType),
			Value:           category.TotalSpent,
		}
	}

	creditCards := make([]dtos.CreditCardsSpendingResponse, len(spendingCreditCards))

	for i, creditCard := range spendingCreditCards {
		creditCards[i] = dtos.CreditCardsSpendingResponse{
			ID:               creditCard.ID,
			Name:             creditCard.Name,
			FirstFourNumbers: creditCard.FirstFourNumbers,
			Limit:            creditCard.Limit,
			CloseDay:         creditCard.CloseDay,
			ExpireDay:        creditCard.ExpireDay,
			BackgroundColor:  creditCard.BackgroundColor,
			TextColor:        creditCard.TextColor,
			TotalSpent:       creditCard.TotalSpent,
		}
	}

	var mostSpentCategory *dtos.CategoriesSpendingResponse

	if len(categories) > 0 {
		mostSpentCategory = &categories[0]
	}

	var mostSpentCreditCard *dtos.CreditCardsSpendingResponse

	if len(creditCards) > 0 {
		mostSpentCreditCard = &creditCards[0]
	}

	return dtos.MonthlyReportResponse{
		TotalIncome:         model.TotalIncome,
		TotalDebit:          model.TotalDebit,
		TotalCredit:         model.TotalCredit,
		Balance:             model.Balance,
		MostSpentCategory:   mostSpentCategory,
		MostSpentCreditCard: mostSpentCreditCard,
		Categories:          categories,
		CreditCards:         creditCards,
	}
}
