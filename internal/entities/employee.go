package entities

import (
	"github.com/marcioc0sta/finch/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID            entities.ID `json:"id"`
	Name          string      `json:"name"`
	Role          string      `json:"role"`
	Email         string      `json:"email"`
	Mobile        string      `json:"mobile"`
	Team          string      `json:"team"`
	Department    string      `json:"department"`
	Address       string      `json:"address"`
	DirectManager string      `json:"direct_manager"`
	Password      string      `json:"password"`
}

func NewEmployee(name, role, email, mobile, team, department, address, directManager, password string) (*Employee, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Employee{
		Name:          name,
		Role:          role,
		Email:         email,
		Mobile:        mobile,
		Team:          team,
		Department:    department,
		Address:       address,
		DirectManager: directManager,
		Password:      string(hash),
	}, nil
}

func (e *Employee) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err == nil
}
