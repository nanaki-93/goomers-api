package repository

import (
	"goomers-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

func NewUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func (repo *PostgresUserRepository) GetByID(id string) (*model.User, error) {
	user := new(model.User)
	repo.db.First(&user, "id = ?", id)
	return user, nil
}

func (repo *PostgresUserRepository) GetAll() ([]*model.User, error) {
	users := make([]*model.User, 0)
	repo.db.Find(&users)
	return users, nil
}

func (repo *PostgresUserRepository) Create(user *model.User) error {
	repo.db.Create(user)
	return nil
}
func (repo *PostgresUserRepository) Update(user *model.User) error {
	repo.db.Save(user)
	return nil
}

func (repo *PostgresUserRepository) Delete(id int) error {
	repo.db.Delete(model.User{}, id)
	return nil
}
