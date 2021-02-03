package databaseLayer

import "wancloudsV2/gen/models"

//UpdateUserInDb
func UpdateUserInDb(paramsUpdateData models.UpdateUsersDefinition)bool{
	_,err := DatabaseConnection()
	if err!=nil{
		return false
	}
	return true


}
