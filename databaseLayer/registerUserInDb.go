package databaseLayer

import (
	"fmt"
	"github.com/taskAPi/gen/models"
)

func RegisterUserInDb(paramsRegisterData models.RegisterUserDefinition) error {
	db, err := DatabaseConnection()
	if err != nil {
		return err
	}
	type data models.RegisterUserDefinition
	db.AutoMigrate(&data{})
	db.Create(data{Email: paramsRegisterData.Email,Name: paramsRegisterData.Name,Password: paramsRegisterData.Password})

	return nil
}

func VerifyDataRedundancy(paramsRegisterData models.RegisterUserDefinition)bool{
	db, err := DatabaseConnection()
	if err!=nil{
		return false
	}

	type Result struct {
		Email string
	}

	var result Result
	db.Table("data").Select("Email").Where("Email = ?", paramsRegisterData.Email).Scan(&result)

	fmt.Println(result.Email)
	fmt.Println(*paramsRegisterData.Email)

	if *paramsRegisterData.Email==result.Email{
		fmt.Print("there is a data redundancy")
		return false
	}
	fmt.Print("there is no data redundancy")
	return true
}
