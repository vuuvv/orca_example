package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vuuvv/orca-example/controllers"
	"github.com/vuuvv/orca/server"
)
import "github.com/vuuvv/orca"

func main() {
	app := orca.NewApplication()
	authMiddleware := server.MiddlewareJwt(
		app.GetHttpServer().GetConfig(),
		server.SimpleAuthorization{
			AnonymousRoutes: map[string]bool{
				"/user":        true,
				"/user/login":        true,
				"/user/login/errors": true,
				"/user/_init_":       true,
			},
		},
	)
	app.Use(server.MiddlewareId, gin.Logger(), gin.Recovery(), authMiddleware).Mount(
		controllers.UserController,
	).Start()
}
