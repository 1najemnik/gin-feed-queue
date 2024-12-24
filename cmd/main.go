package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"gin-feed-queue/controllers"
	"gin-feed-queue/middlewares"
	"gin-feed-queue/services"
)

func main() {
	firestoreClient := initFirebase()
	defer firestoreClient.Close()

	services.SetFirestoreClient(firestoreClient)
	services.InitTelegramBot()

	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	r.Use(middlewares.ValidateAccessKey())
	r.SetFuncMap(map[string]interface{}{
		"HasStatus":        controllers.HasStatus,
		"GetStatusStrings": controllers.GetStatusStrings,
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.RenderIndexPage)
	r.GET("/news/edit/:id", controllers.RenderEditNewsPage)
	r.POST("/news/edit/:id", controllers.EditNews)
	r.POST("/api/rss", controllers.FetchRSSNews)
	r.POST("/news/publish/:id", controllers.PublishToTelegram)
	r.POST("/news/delete/:id", controllers.DeleteNews)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}

func initFirebase() *firestore.Client {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	credentialsPath := filepath.Join(currentDir, "config", "serviceAccountKey.json")

	opt := option.WithCredentialsFile(credentialsPath)

	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DATABASE_URL"),
	}, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
	}

	return client
}
