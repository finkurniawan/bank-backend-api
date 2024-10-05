package service

import (
	"fmt"

	_ "github.com/finkurniawan/bank-backend-api/api/model"
	"github.com/finkurniawan/bank-backend-api/api/repository"
	"github.com/finkurniawan/bank-backend-api/api/utils"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	customerRepository repository.CustomerRepository
}

func NewAuthService(customerRepository repository.CustomerRepository) AuthService {
	return &authService{customerRepository: customerRepository}
}

func (s *authService) Login(username, password string) (string, error) {
	customer, err := s.customerRepository.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, customer.Password) {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := utils.GenerateToken(fmt.Sprint(customer.ID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}
