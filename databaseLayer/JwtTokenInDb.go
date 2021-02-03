package databaseLayer

import "wancloudsV2/Models"

func SaveJwtToken(paramsForJwt Models.JwtModels){
	db, err:= DatabaseConnection()
	if err!=nil{
		err.Error()
	}
	type Tokens Models.JwtModels

	db.AutoMigrate(&Tokens{})
	db.Create(&Tokens{Token: paramsForJwt.Token,Email: paramsForJwt.Email })
}

