package services

import (
	"wancloudsV2/databaseLayer"
	"wancloudsV2/gen/models"
)

func(s *Service) LoginUserService(paramsLoginData models.LoginUserDefinition) bool{
	loginSuccessful:=databaseLayer.LoginUserInDb(paramsLoginData)
	if loginSuccessful == true{
		return true
	}
	return false
}