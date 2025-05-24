package domain

import "time"

type Category struct {
	ID *string
	Name string
	UserID *string
	CreatedAt string
}


type SubCategory struct {
	ID *string
	Name string
	CategoryID *string
	UserID *string
	CreatedAt string
}	


func NewCategory(
	name string,
	userID *string,
) *Category {
	return &Category{
		Name:     name,
		CreatedAt:  time.Now().Format(time.DateTime),
		UserID:    userID,
	}
}

func NewSubCategory(
	name string,
	categoryID *string,
	userID *string,
) *SubCategory {
	return &SubCategory{
		Name:      name,
		CategoryID: categoryID,
		UserID:    userID,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}
