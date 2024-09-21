package main

import (
	"net/http"
	"time"

	"github.com/Anrewp/go_balance/src/config"
	"github.com/Anrewp/go_balance/src/controller"
	"github.com/Anrewp/go_balance/src/helper"
	"github.com/Anrewp/go_balance/src/model"
	"github.com/Anrewp/go_balance/src/repository"
	"github.com/Anrewp/go_balance/src/router"
	"github.com/Anrewp/go_balance/src/service"
)

// todo
func main() {
	db := config.DatabaseConnection()
	db.Table("users").AutoMigrate(&model.Users{})
	userRepository := repository.NewUsersRepositoryImpl(db)
	userService := service.NewUsersServiceImpl(userRepository)
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
