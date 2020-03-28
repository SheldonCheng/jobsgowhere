package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jobsgowhere/jobsgowhere/api/controllers"
	"github.com/jobsgowhere/jobsgowhere/api/middlewares"
	"github.com/jobsgowhere/jobsgowhere/api/util"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	taskController := controllers.TaskController{}
	taskController.Init(db)

	router.Use(func(ctx *gin.Context) {
		if !util.Contains([]string{"POST", "PUT", "PATCH"}, ctx.Request.Method) {
			return
		}

		if ctx.Request.Header["Content-Length"][0] == "0" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Payload should not be empty"})
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if len(ctx.Request.Header["Content-Type"]) == 0 ||
			!util.Contains(ctx.Request.Header["Content-Type"], "application/json") {
			ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"message": "Content type should be application/json"})
			ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}
	})

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0ClientID := os.Getenv("AUTH0_CLIENT_ID")
	auth0Audience := os.Getenv("AUTH0_AUDIENCE")
	auth0Callback := os.Getenv("AUTH0_CALLBACK")

	if auth0Domain == "" || auth0ClientID == "" || auth0Audience == "" || auth0Callback == "" {
		log.Panic("AUTH0 details not found in environment variables")
	}

	dataToUIPage := gin.H{
		"AUTH0_DOMAIN":    auth0Domain,
		"AUTH0_CLIENT_ID": auth0ClientID,
		"AUTH0_AUDIENCE":  auth0Audience,
		"AUTH0_CALLBACK":  auth0Callback,
	}

	router.LoadHTMLGlob("ui-dist/*.html")
	router.Static("/static", "./ui-dist/static")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", dataToUIPage)
	})

	router.POST("/api/tasks", middlewares.AuthMiddleware(), taskController.CreateTask)
	router.GET("/api/tasks", middlewares.AuthMiddleware(), taskController.GetTasks)
	router.GET("/api/tasks/:id", middlewares.AuthMiddleware(), taskController.GetTaskByID)
	router.PUT("/api/tasks/:id", middlewares.AuthMiddleware(), taskController.UpdateTaskForID)
	router.DELETE("/api/tasks/:id", middlewares.AuthMiddleware(), taskController.DeleteTaskForID)


	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", dataToUIPage)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)
}
