package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthEndpoints(router *gin.Engine) {
	router.GET("/health", func (c *gin.Context) {
		fmt.Println("Health check called!")

		c.IndentedJSON(http.StatusOK, "Healthy")
	})
}