package main

import (
	"fmt"
	"go-echo-clean-architecuture/db"
	"go-echo-clean-architecuture/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
