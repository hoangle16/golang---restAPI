package user

import (
	"restful_api/entities"

	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	DBConnection *gorm.DB
}

// NewSqliteUserRepository - return UserRepository
func NewSqliteUserRepository(DBConn *gorm.DB) Repository {
	return &sqliteUserRepository{DBConnection: DBConn}
}

func (s *sqliteUserRepository) FetchAll() ([]*entities.User, error) {
	users := make([]*entities.User, 0)
	if err := s.DBConnection.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *sqliteUserRepository) FindByID(id uint) (*entities.User, error) {
	user := entities.User{}
	if err := s.DBConnection.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *sqliteUserRepository) Store(u *entities.User) (bool, error) {
	if err := s.DBConnection.Create(&u).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *sqliteUserRepository) Update(u *entities.User) (bool, error) {
	if err := s.DBConnection.Save(&u).Error; err != nil {
		return false, nil
	}
	return true, nil
}
func (s *sqliteUserRepository) Delete(id uint) (bool, error) {
	if err := s.DBConnection.Delete(entities.User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
