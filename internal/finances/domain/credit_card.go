package domain


type CreditCard struct {
	ID *string
	Name string
	LimitTotal float64
	LimitAvailable float64
	AccountID *string
	UserID *string
	CreatedAt string
}

func NewCreditCard(
	name string,
	limitTotal float64,
	limitAvailable float64,
	accountID *string,
	userID *string,
) *CreditCard {
	return &CreditCard{
		Name:            name,
		LimitTotal:      limitTotal,
		LimitAvailable:  limitAvailable,
		AccountID:       accountID,
		UserID:          userID,
	}
}

