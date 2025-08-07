package service

import "main/internal/core/port"

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(
	userRepo port.UserRepository,
) *UserService {
	return &UserService{
		userRepo,
	}
}
