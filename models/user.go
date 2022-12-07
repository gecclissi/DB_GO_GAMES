package models

type Users struct {
	Userid int    `json:"userId" gorm:"primaryKey"`
	Email  string `json:"email"`
}

type UsersInput struct {
	Email string `json:"email" binding:"required"`
}
