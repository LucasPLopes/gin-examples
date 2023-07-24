package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db = make(map[string]string)

type User struct {
	User   string `json:"user"`
	Info   string `json:"info"`
	Status string `json:"status"`
}

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		info, ok := db[user]
		if ok {
			c.JSON(http.StatusOK,
				User{
					User: user,
					Info: info,
				})
		} else {
			c.JSON(http.StatusOK,
				User{
					User:   user,
					Status: "no info",
				})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":   "bar",      // Zm9vOmJhcg==
		"admin": "password", // YWRtaW46cGFzc3dvcmQ=
	}))

	authorized.POST("admin", func(c *gin.Context) {
		// user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			User string `json:"user" binding:"required"`
			Info string `json:"info" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[json.User] = json.Info
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := Router()
	r.Run(fmt.Sprintf(":%s", port))
}
