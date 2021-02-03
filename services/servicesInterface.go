package services

import (
	"gorm.io/gorm"
)

type Service struct {
	dbConnection *gorm.DB
}

func NewService(dbConnection *gorm.DB) *Service {
	return &Service{dbConnection: dbConnection}
}


