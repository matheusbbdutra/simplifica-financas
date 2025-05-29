package dto

type UpdateUserRequest struct {
	Name *string `json:"name"`
	Email    *string `json:"email" validate:"email"`
	Password *string `json:"password" validate:"min=8"`
}