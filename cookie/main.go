package main

import (
	"net/http"

	ginexamples "github.com/LucasPLopes/gin-examples"
	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden no cookie"})
		c.Abort()

	}
}

func main() {
	route := gin.Default()
	route.GET("/login", func(c *gin.Context) {
		c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
		c.String(200, "Login success")
	})

	route.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Your home page"})
	})

	route.Run(ginexamples.GetPort())
}
