package service

import (
	"time"

	"go-projects/internal/models"
)

type UserService struct {
	users  []models.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  []models.User{},
		nextID: 1,
	}
}

// CREATE USER
func (s *UserService) CreateUser(name string, dob string) models.User {
	parsedDob, _ := time.Parse("2006-01-02", dob)

	user := models.User{
		ID:   s.nextID,
		Name: name,
		Dob:  dob,
		Age:  CalculateAge(parsedDob),
	}

	s.users = append(s.users, user)
	s.nextID++

	return user
}

// GET USER
func (s *UserService) GetUserByID(id int) (models.User, bool) {
	for _, user := range s.users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

// LIST USERS
func (s *UserService) GetAllUsers() []models.User {
	return s.users
}
