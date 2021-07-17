package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"possant-admin/internal/domain/user/entity"
)

func Database() *gorm.DB{
	user := "quiszzibrvczyp"
	pass := "508f19a87738809dcac86a97353d430193d7b050a1147fa3fd1504f95cf9e5a3"
	database := "datcn8gkm40bds"
	port := "5432"
	host := "ec2-52-7-115-250.compute-1.amazonaws.com"

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + database + " port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	return db
}
