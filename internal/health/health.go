package health

import (
	"context"
	"fmt"
	"net/http"

	"narrativeboard/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func RegisterHealthEndpoints(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		fmt.Println("Health check called!")

		c.IndentedJSON(http.StatusOK, "Healthy")
	})

	router.GET("/health/dbcheck", func(c *gin.Context) {
		result, err := database.Run(func(conn *pgx.Conn) (result any, err error) {
			row := conn.QueryRow(context.Background(), "select 'Hello!'")
			err = row.Scan(&result)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return nil, err
			}
			return
		})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, fmt.Sprintf("Error!: %v", result))
			return
		}

		fmt.Printf("result: %v\n", result)
		c.IndentedJSON(http.StatusOK, "Success")
	})
}
