package database

import (
	"github.com/marcioc0sta/finch/internal/entities"
	"gorm.io/gorm"
)

type Employee struct {
	DB *gorm.DB
}

func NewEmployee(db *gorm.DB) *Employee {
	return &Employee{DB: db}
}

func (e *Employee) Create(employee *entities.Employee) error {
	return e.DB.Create(employee).Error
}

func (e *Employee) FindByEmail(email string) (*entities.Employee, error) {
	var employee entities.Employee
	err := e.DB.Where("email = ?", email).Find(&employee).Error

	return &employee, err
}
