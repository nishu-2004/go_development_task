package service

import (
	"context"

	"go-projects/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CREATE USER
func (s *UserService) CreateUser(ctx context.Context, name string, dob string) (map[string]interface{}, error) {
	user, err := s.repo.CreateUser(ctx, name, dob)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  CalculateAge(user.Dob.Time),
	}, nil
}

// GET USER BY ID
func (s *UserService) GetUserByID(ctx context.Context, id int32) (map[string]interface{}, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  CalculateAge(user.Dob.Time),
	}, nil
}

// LIST USERS
func (s *UserService) ListUsers(ctx context.Context) ([]map[string]interface{}, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, u := range users {
		result = append(result, map[string]interface{}{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Time.Format("2006-01-02"),
			"age":  CalculateAge(u.Dob.Time),
		})
	}

	return result, nil
}
