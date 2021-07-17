//+build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	user "possant-admin/internal/domain/user/handler"
	"possant-admin/internal/domain/user/repository"
	"possant-admin/internal/domain/user/service"
)

func InjectUser(db *gorm.DB) user.UserHandler {
	wire.Build(repository.UserRepositoryProvider, service.UserServiceProvider, user.UserHandlerProvider)
	return user.UserHandler{}
}