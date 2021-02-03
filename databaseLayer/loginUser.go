package databaseLayer

import (
	"fmt"
	"wancloudsV2/gen/models"
	"wancloudsV2/Models"
)

func LoginUserInDb(paramsLoginData models.LoginUserDefinition) bool {
	db, err := DatabaseConnection()
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	results:= Models.Results{}
	db.Table("data").Select("Email,Password").Where("Email = ?", paramsLoginData.Email).Scan(&results)
	if *paramsLoginData.Email!=results.Email && *paramsLoginData.Password != results.Password{
		return false
	}
	return true
}
