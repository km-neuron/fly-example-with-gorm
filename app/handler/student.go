package handler

import (
	"github.com/rg-app/app/model"

	"gorm.io/gorm"
)

type StudentHandler struct {
	db *gorm.DB
}

func NewStudentHandler(db *gorm.DB) StudentHandler {
	return StudentHandler{db}
}

func (u *StudentHandler) AddStudent(student model.Student) error {
	if result := u.db.Create(&student); result.Error != nil {
		return gorm.ErrInvalidData
	}

	return nil
}

func (u *StudentHandler) ReadStudent() ([]model.Student, error) {
	rows, err := u.db.Table("students").Where("deleted_at IS NULL").Select("id, username, grade").Rows()
	if err != nil {
		return nil, err
	}

	var listTodo []model.Student
	for rows.Next() {
		u.db.ScanRows(rows, &listTodo)
	}

	return listTodo, nil
}
