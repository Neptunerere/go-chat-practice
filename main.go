package main

import (
	"chat-server/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  utils.BUFFER_SIZE * utils.BUFFER_SIZE * utils.BUFFER_SIZE,
	WriteBufferSize: utils.BUFFER_SIZE * utils.BUFFER_SIZE * utils.BUFFER_SIZE,

	// Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func initDb() string {
	config :=
		utils.DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Db:       os.Getenv("DB_DATABASE"),
		}

	return config.GetConnConfigs()
}

func setupRouter() *gin.Engine {
	// gin init
	r := gin.New()

	r.Use(gin.Logger())   // Logger set
	r.Use(gin.Recovery()) // middleware Panic 500 Error set

	api := r.Group("/chat")

	// Setting up a router group and adding middleware for "/chat/v1".
	v1 := api.Group("/v1")
	v1.Use(func(c *gin.Context) {
		c.Set("Version", "v1")
		c.Next()
	})

	return r
}

func main() {
	// ... (Setting up routers and middleware, etc.)
	r := setupRouter()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database setting
	connectionStr := initDb()

	connector, err := utils.Connect(connectionStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(connector)

	// Server start
	serverErr := r.Run(fmt.Sprintf(":%v", os.Getenv("CHAT_SERVER_PORT")))
	if serverErr != nil {
		log.Fatalf("Failed to start the server: %v", serverErr)
	}
}
