package models

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
