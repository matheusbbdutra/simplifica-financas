package domain

import "time"

type TransactionType string

const (
	Income TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	ID *string
	Description string
	Amount float64
	DueDate string
	PaidAt *string
	IsEffective bool
	CategoryID *string
	SubCategoryID *string
	AccountID *string
	CreditCardID *string
	IsRecurring bool
	InstallmentNumber *int
	InstallmentTotal *int
	TransactionTypeID *TransactionType
	UserID *string
	CreatedAt string
}

func NewTransaction(
	description string,
	amount float64, 
	dueDate string, 
	paidAt *string, 
	isEffective bool, 
	categoryID *string, 
	subCategoryID *string, 
	accountID *string, 
	creditCardID *string, 
	isRecurring bool, 
	installmentNumber *int, 
	installmentTotal *int, 
	transactionTypeID *TransactionType, 
	userID *string,
) *Transaction {
	return &Transaction{
		Description:      description,
		Amount:           amount,
		DueDate:          dueDate,
		PaidAt:           paidAt,
		IsEffective:      isEffective,
		CategoryID:       categoryID,
		SubCategoryID:    subCategoryID,
		AccountID:        accountID,
		CreditCardID:     creditCardID,
		IsRecurring:      isRecurring,
		InstallmentNumber: installmentNumber,
		InstallmentTotal: installmentTotal,
		TransactionTypeID: transactionTypeID,
		UserID:           userID,
		CreatedAt:        time.Now().Format(time.DateTime),
	}
}