package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnvDefault(key string, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defVal
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": 200,
			"message":    "Hello World!",
		})
	})

	PORT := getEnvDefault("PORT", ":8080")
	r.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
