package services

import (
	"wancloudsV2/JWT"
	"wancloudsV2/Models"
	"wancloudsV2/databaseLayer"
)

func (s *Service)GeneratingTheJWT(Email string) (string, error){
	Token, err:=JWT.GenerateJWT(Email)
	if err!=nil{
		return "", err
	}
	databaseLayer.SaveJwtToken(Models.JwtModels{Email: Email,Token: Token})
	return Token,nil
}
