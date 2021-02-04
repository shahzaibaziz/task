package databaseLayer

import "github.com/taskAPi/gen/models"

//UpdateUserInDb
func UpdateUserInDb(paramsUpdateData models.UpdateUsersDefinition)bool{
	_,err := DatabaseConnection()
	if err!=nil{
		return false
	}
	return true


}
