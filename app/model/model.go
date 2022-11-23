package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username string `json:"username"`
	Grade    int    `json:"grade"`
}

type StudentView struct {
	Id       uint
	Username string
	Grade    int
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
