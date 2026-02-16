package subjects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func RegisterSubjectsEndpoints(eng *gin.Engine) {
	eng.GET("/subjects", func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Subjects")
	})
}