package subjects

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)


func RegisterSubjectsEndpoints(eng *gin.Engine) {
	eng.GET("/subjects", func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Subjects")
	})

	eng.GET("/subjects/:id", func (c *gin.Context) {
		returnedString := "testing"
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)
		go testing(&returnedString, &waitGroup)
		waitGroup.Wait()
		c.IndentedJSON(http.StatusOK, returnedString)
	})
}
func testing(returnedString *string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	*returnedString = "Even Slower Testing"
}