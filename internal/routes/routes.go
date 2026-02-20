package routes

import (
	"narrativeboard/internal/health"
	"narrativeboard/internal/subjects"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eng *gin.Engine) {
	health.RegisterHealthEndpoints(eng)
	subjects.RegisterSubjectsEndpoints(eng)
}
