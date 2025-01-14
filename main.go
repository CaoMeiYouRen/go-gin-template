package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvDefault(key string, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defVal
}

func loadEnv() {
	// 检查.env文件是否存在
	if _, err := os.Stat(".env"); err == nil {
		// 如果文件存在，加载它
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Println(".env file loaded successfully")
	} else if os.IsNotExist(err) {
		log.Println(".env file not found, skipping load")
	} else {
		log.Fatalf("Error checking .env file: %v", err)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": 200,
			"message":    "Hello World!",
		})
	})
	return r
}

func main() {
	loadEnv()
	r := SetupRouter()
	port := getEnvDefault("PORT", "8080")
	host := getEnvDefault("HOST", "0.0.0.0")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	addr := host + ":" + port
	r.Run(addr)

}
