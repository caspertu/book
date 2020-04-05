package main

import (
	"github.com/caspertu/book/api/v1/books"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	r := gin.Default()
	books.SetupRouter(r)
	log.Info("Server started")
	log.Fatalln(r.Run())
}