package main

import (
	"blog-backend/controllers"
	"blog-backend/database"
	"blog-backend/diagnostics"
	"blog-backend/docs"
	"blog-backend/middleware"
	"blog-backend/util"
	"cloud.google.com/go/firestore"
	"embed"
	"flag"
	"fmt"
	armor "github.com/bmwadforth-com/armor-go/src"
	armorUtil "github.com/bmwadforth-com/armor-go/src/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	port = flag.Int("port", 8080, "The server port")

	//go:embed web/build
	web embed.FS

	logCleanup func(*zap.Logger)

	r *gin.Engine
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

func init() {
	var err error
	err = util.SetupArmor()
	if err != nil {
		panic(err)
	}

	r = gin.New()
	if util.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
	}

	database.DbConnection, err = firestore.NewClientWithDatabase(armor.ArmorContext, util.Config.ProjectId, util.Config.FireStoreDatabase)
	if err != nil {
		armorUtil.SLogger.Fatalf("failed to create firestore client: %v", err)
	}

	prometheus.MustRegister(diagnostics.ArticlesCounter)
}

func main() {
	defer armor.CleanupLogger(armorUtil.Logger)

	flag.Parse()
	defer func(DbConnection *firestore.Client) {
		err := DbConnection.Close()
		if err != nil {
			armorUtil.SLogger.Fatalf("failed to close db client: %v", err)
		}
	}(database.DbConnection)

	wwwroot := EmbedFolder(web, "web/build")
	staticServer := static.Serve("/", wwwroot)

	docs.SwaggerInfo.BasePath = "/api/v1"

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

	if !util.IsProduction {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	v1 := r.Group("/api/v1")
	v1.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1.GET("/ping", controllers.Ping)
	v1.GET("/healthz", controllers.HealthCheck)
	v1.GET("/articles", controllers.GetArticles)
	v1.GET("/article/:slug", controllers.GetArticleBySlug)
	v1.POST("/login", controllers.LoginUser)

	v1ApiKeyAuthenticated := r.Group("/api/v1")
	v1ApiKeyAuthenticated.Use(middleware.ApiKeyAuthenticationMiddleware())
	v1ApiKeyAuthenticated.POST("/user", controllers.CreateUser)

	v1BearerAuthenticated := r.Group("/api/v1")
	v1BearerAuthenticated.Use(middleware.BearerAuthenticationMiddleware())
	v1BearerAuthenticated.POST("/article", controllers.CreateArticle)
	v1BearerAuthenticated.POST("/article/:articleId/content", controllers.UploadArticleContent)
	v1BearerAuthenticated.GET("/sessions", controllers.GetSessions)
	v1BearerAuthenticated.GET("/gemini", controllers.QueryGemini)
	v1BearerAuthenticated.GET("/status", controllers.GetStatus)

	err := r.SetTrustedProxies([]string{})
	if err != nil {
		armorUtil.SLogger.Fatalf("an error has occurred: %v", err)
	}

	err = r.Run(fmt.Sprintf(":%s", strconv.Itoa(*port)))
	if err != nil {
		armorUtil.SLogger.Fatalf("unable to start blog-backend: %v", err)
	}
}
