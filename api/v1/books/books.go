package books

import (
	"github.com/caspertu/book/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	apiv1books := r.Group("/api/v1/books")
	{
		apiv1books.GET("/", fetchAll)
		apiv1books.GET("/:id", fetchSingle)
		apiv1books.POST("/", create)
		apiv1books.DELETE("/:id", remove)
		apiv1books.PUT("/:id", update)
	}
	return r
}

// GET /api/v1/books/
func fetchAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.FindAll()})
}

// fetchSingle 获得一个
// GET /api/v1/books/:ID
func fetchSingle(c *gin.Context) {
	param := c.Param("id")
	ID, _ := strconv.Atoi(param)
	book, err := model.FindBook(ID)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// create ...
// POST /api/v1/books/
func create(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	book.Create()
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /api/v1/books/:id
func remove(c *gin.Context) {
	param := c.Param("id")
	ID, _ := strconv.Atoi(param)
	book, err := model.FindBook(ID)
	if err != nil {
		panic(err)
	}
	book.Delete()
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UPDATE /api/v1/books/:id
func update(c *gin.Context) {
	param := c.Param("id")
	ID, _ := strconv.Atoi(param)
	book, err := model.FindBook(ID)
	if err != nil {
		panic(err)
	}
	book.Title = "GoGo"
	book.Update()
	c.JSON(http.StatusOK, gin.H{"data": book})
}
