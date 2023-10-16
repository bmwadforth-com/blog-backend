package main

import (
	"blog-backend/controllers"
	"blog-backend/middleware"
	"blog-backend/util"
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 8080, "The server port")

	//go:embed web/build
	web embed.FS
)

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

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

	wwwroot := EmbedFolder(web, "web/build")
	staticServer := static.Serve("/", wwwroot)

	r.Use(gin.Recovery())
	r.Use(staticServer)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	v1 := r.Group("/api/v1")
	v1.GET("/ping", controllers.Ping)
	v1.GET("/healthz", controllers.HealthCheck)

	v1.POST("/article", controllers.CreateArticle)
	v1.POST("/article/:articleId/content", controllers.UploadArticle)
	v1.GET("/articles", controllers.GetArticles)

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
		util.SLogger.Fatalf("unable to start blog-backend: %v", err)
	}
}
