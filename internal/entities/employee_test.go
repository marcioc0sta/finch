package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmployee(t *testing.T) {
	user, err := NewEmployee("John Doe", "Developer", "j@j.com", "123456789", "Team 1", "Department 1", "Address 1", "Manager 1", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "Developer", user.Role)
	assert.Equal(t, "j@j.com", user.Email)
	assert.Equal(t, "123456789", user.Mobile)
	assert.Equal(t, "Team 1", user.Team)
	assert.Equal(t, "Department 1", user.Department)
	assert.Equal(t, "Address 1", user.Address)
	assert.Equal(t, "Manager 1", user.DirectManager)
	assert.Equal(t, "$2a$10$", user.Password[:7])
}

func TestValidatePassword(t *testing.T) {
	user, err := NewEmployee("John Doe", "Developer", "j@j.com", "123456789", "Team 1", "Department 1", "Address 1", "Manager 1", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
