package main

import (
	"github.com/caspertu/book/api/v1/books"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	books.SetupRouter(r)
	r.Run()
}