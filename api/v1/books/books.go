package books

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) (*gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	apiv1books := r.Group("/api/v1/books")
	{
		apiv1books.GET("/:id", FetchSingle)
	}
	return r
}

// :ID
func FetchSingle(c *gin.Context) {
	ID := c.Param("id")
	//data := gin.H{
	//	"data": ID,
	//}
	//c.IndentedJSON(http.StatusOK, data)
	c.String(http.StatusOK, ID)
}
