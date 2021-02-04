package services

import (
	"github.com/taskAPi/databaseLayer"
	"github.com/taskAPi/gen/models"
)

func(s *Service) LoginUserService(paramsLoginData models.LoginUserDefinition) bool{
	loginSuccessful:=databaseLayer.LoginUserInDb(paramsLoginData)
	if loginSuccessful == true{
		return true
	}
	return false
}