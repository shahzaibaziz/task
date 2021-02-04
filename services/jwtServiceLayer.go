package services

import (
	"github.com/taskAPi/JWT"
	"github.com/taskAPi/Models"
	"github.com/taskAPi/databaseLayer"
)

func (s *Service)GeneratingTheJWT(Email string) (string, error){
	Token, err:=JWT.GenerateJWT(Email)
	if err!=nil{
		return "", err
	}
	databaseLayer.SaveJwtToken(Models.JwtModels{Email: Email,Token: Token})
	return Token,nil
}
