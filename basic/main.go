package main

import (
	"net/http"

	ginexamples "github.com/LucasPLopes/gin-examples"
	"github.com/gin-gonic/gin"
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
	port := ginexamples.GetPort()
	r := Router()
	r.Run(port)
}
