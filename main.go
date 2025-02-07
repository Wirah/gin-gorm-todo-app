package main

import (
	"fmt"

	"github.com/Wirah/gin-gorm-todo-app/Config"
	"github.com/Wirah/gin-gorm-todo-app/Models"
	"github.com/Wirah/gin-gorm-todo-app/Routes"

	"github.com/go-gorm/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
