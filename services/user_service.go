package services

import (
	"errors"

	"github.com/KevinMaulanaAtmaja/project-management-golang/models"
	"github.com/KevinMaulanaAtmaja/project-management-golang/repositories"
	"github.com/KevinMaulanaAtmaja/project-management-golang/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	// harus cek email yg terdaftar apkah sudah dipake/blum
	// hashing password
	// set role
	// simpan user

	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("Email already registered")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid Credential")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid Credential")
	}

	return user, nil
}
