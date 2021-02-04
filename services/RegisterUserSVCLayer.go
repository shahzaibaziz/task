package services

import (
	"github.com/taskAPi/databaseLayer"
	"github.com/taskAPi/gen/models"
)

func (s *Service)VerifyData(paramsVerifyData models.RegisterUserDefinition)bool{
	dataVerified :=databaseLayer.VerifyDataRedundancy(paramsVerifyData)
	if dataVerified==false{
		return false
	}
	return true
}

func (s *Service)RegisterUserService(paramsRegisterData models.RegisterUserDefinition)  {

	_ = databaseLayer.RegisterUserInDb(paramsRegisterData)

}
