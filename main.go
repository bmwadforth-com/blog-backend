package main

import (
	"flag"
	"os"
	"web-template/controllers"
	"web-template/middleware"
	"web-template/util"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	logCleanup := util.InitLogger()
	defer logCleanup(util.Logger)

	r := gin.New()

	util.IsProduction = os.Getenv("APP_ENV") == "PRODUCTION"
	if util.IsProduction {
		util.LoadEnvironmentVariables()
		gin.SetMode(gin.ReleaseMode)
	} else {
		util.LoadConfiguration()
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.GET("/ping", controllers.Ping)
	v1.GET("/healthz", controllers.HealthCheck)

	v1.POST("/person", controllers.CreatePerson)
	v1.GET("/person", controllers.GetPeople)

	v1ApiKeyAuthenticated := r.Group("/api/v1")
	v1ApiKeyAuthenticated.Use(middleware.ApiKeyAuthenticationMiddleware())

	v1BearerAuthenticated := r.Group("/api/v1")
	v1BearerAuthenticated.Use(middleware.BearerAuthenticationMiddleware())

	err := r.SetTrustedProxies([]string{})
	if err != nil {
		util.SLogger.Fatalf("an error has occurred: %v", err)
	}

	err = r.Run()
	if err != nil {
		util.SLogger.Fatalf("unable to start web-template: %v", err)
	}
}
