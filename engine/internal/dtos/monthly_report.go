package dtos

import (
	"github.com/google/uuid"
)

type MonthlyReportResponse struct {
	TotalIncome         float64                       `json:"total_income"`
	TotalDebit          float64                       `json:"total_debit"`
	TotalCredit         float64                       `json:"total_credit"`
	Balance             float64                       `json:"balance"`
	MostSpentCategory   *CategoriesSpendingResponse   `json:"most_spent_category,omitempty"`
	MostSpentCreditCard *CreditCardsSpendingResponse  `json:"most_spent_creditcard,omitempty"`
	Categories          []CategoriesSpendingResponse  `json:"categories"`
	CreditCards         []CreditCardsSpendingResponse `json:"creditcards"`
}

type CategoriesSpendingResponse struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Icon            string    `json:"icon"`
	TransactionType int       `json:"transaction_type"`
	Value           float64   `json:"value"`
}

type CreditCardsSpendingResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
	BackgroundColor  string    `json:"background_color"`
	TextColor        string    `json:"text_color"`
	TotalSpent       float64   `json:"total_spent"`
}
