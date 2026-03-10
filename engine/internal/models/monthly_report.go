package models

import "github.com/google/uuid"

type MonthlyReport struct {
	TotalIncome float64
	TotalDebit  float64
	TotalCredit float64
	Balance     float64
}

type CategoriesSpending struct {
	CategoryID              uuid.UUID
	CategoryName            string
	CategoryIcon            string
	CategoryTransactionType TransactionType
	TotalSpent              float64
}

type CreditCardsSpending struct {
	ID               uuid.UUID
	Name             string
	FirstFourNumbers string
	Limit            float64
	CloseDay         int32
	ExpireDay        int32
	BackgroundColor  string
	TextColor        string
	TotalSpent       float64
}
