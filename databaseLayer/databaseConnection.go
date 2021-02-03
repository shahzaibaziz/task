package databaseLayer

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnection() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/taskDb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("there is an error in the database connection ")
		//return users.NewRegisterUserBadRequest()
	}
	fmt.Println("you are connected")
	return db, nil
}
