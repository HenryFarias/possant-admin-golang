package repository

import (
	"gorm.io/gorm"
	"possant-admin/internal/domain/user/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryProvider(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) FindAll() []entity.User {
	var users []entity.User
	r.db.Find(&users)
	return users
}

func (r *UserRepository) Save(user entity.User) {
	r.db.Save(&user)
}
