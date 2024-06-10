package main

import (
	"github.com/Mubinabd/auth_service/api"
	"github.com/Mubinabd/auth_service/api/handlers"
	"github.com/Mubinabd/auth_service/service"
	"github.com/Mubinabd/auth_service/storage/postgres"
)


func main() {
	db,err := postgres.ConnectDB()
	if err!=nil{
        panic(err)
    }

	userService := service.NewUserService(db)
	h := handlers.NewHandler(userService)

	router := api.NewGin(h)
	router.Run(":8080")
}