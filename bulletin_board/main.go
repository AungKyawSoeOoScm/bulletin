package main

import (
	"gin_test/bulletin_board/controller"
	postinterfaces "gin_test/bulletin_board/dao/post"
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/helper"
	config "gin_test/bulletin_board/initializers"
	"gin_test/bulletin_board/router"
	authService "gin_test/bulletin_board/service/auth"
	postService "gin_test/bulletin_board/service/post"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDatabase()
}
func main() {
	log.Info().Msg("Server Started")

	// db := config.ConnectDatabase()
	validate := validator.New()
	// db.Table("posts").AutoMigrate(&model.Posts{})
	// db.Table("users").AutoMigrate(&model.User{})

	// Users
	userInterface := interfaces.NewUsersInterfaceImpl(config.DB)
	authService := authService.NewAuthServiceImpl(userInterface, validate)
	authController := controller.NewAuthController(authService)
	userController := controller.NewUsercontroller(userInterface)

	// Posts
	postsInterface := postinterfaces.NewPostsRepositoryImpl(config.DB)
	postsService := postService.NewPostsRepositoryImpl(postsInterface, validate)
	postsController := controller.NewPostsController(postsService)

	routes := router.NewRouter(authController, userController, postsController, userInterface)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
