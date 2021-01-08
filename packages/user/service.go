package user

import (
	"restful_api/entities"

	"golang.org/x/crypto/bcrypt"
)

// Service struct
type Service struct {
	repo Repository
}

// NewService - return &Service
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// FetchAll implement
func (s *Service) FetchAll() ([]*entities.User, error) {
	return s.repo.FetchAll()
}

// FindByID implement
func (s *Service) FindByID(id uint) (*entities.User, error) {
	return s.repo.FindByID(id)
}

// Store implement
func (s *Service) Store(u *entities.User) (bool, error) {
	passHash, err := s.HashPassword(string(u.Password))
	if err != nil {
		return false, err
	}
	u.Password = passHash
	return s.repo.Store(u)
}

// Update implement
func (s *Service) Update(u *entities.User) (bool, error) {
	return s.repo.Update(u)
}

// Delete implement
func (s *Service) Delete(id uint) (bool, error) {
	return s.repo.Delete(id)
}

// HashPassword func
func (s *Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}

// CheckPasswordHash func
func (s *Service) CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
