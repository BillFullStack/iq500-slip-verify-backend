package service

import "main/internal/core/port"

type AuthenticationService struct {
	userRepo port.UserRepository
}

func NewAuthenticationService(
	userRepo port.UserRepository,
) *AuthenticationService {
	return &AuthenticationService{
		userRepo,
	}
}
