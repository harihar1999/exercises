package main

import (
	"awesomeProject4/Config"
	"awesomeProject4/Models"
	"awesomeProject4/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(Config.DB)
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	err := r.Run()
	if err != nil {
		return
	}
}