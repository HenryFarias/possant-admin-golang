// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"gorm.io/gorm"
	"possant-admin/internal/domain/user/handler"
	"possant-admin/internal/domain/user/repository"
	"possant-admin/internal/domain/user/service"
)

// Injectors from wire.go:

func InjectUser(db *gorm.DB) handler.UserHandler {
	userRepository := repository.UserRepositoryProvider(db)
	userService := service.UserServiceProvider(userRepository)
	userHandler := handler.UserHandlerProvider(userService)
	return userHandler
}
