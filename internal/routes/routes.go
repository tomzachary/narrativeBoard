package routes

import (
	"narrativereddit.local/health"
	"narrativereddit.local/subjects"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eng *gin.Engine) {
	health.RegisterHealthEndpoints(eng)
	subjects.RegisterSubjectsEndpoints(eng)
}
