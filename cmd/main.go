package main

import (
	"github.com/sally0226/oidc-go-example/api"
	"github.com/sally0226/oidc-go-example/database"
	"github.com/sally0226/oidc-go-example/repository"
	"github.com/sally0226/oidc-go-example/service"
	"os"
)

func main() {
	db := database.MustInitSQLlite()
	database.MustMigrateDB(db)

	userRepository := repository.NewUserRepository(db)
	oauthService := service.NewOAuthServices()
	userService := service.NewUserService(userRepository)
	h := api.NewHandler(oauthService, userService)

	r := api.SetupRouter(h)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
