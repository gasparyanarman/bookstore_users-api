package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// Application entry point
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
