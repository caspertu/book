package main

import (
	"github.com/caspertu/book/api/v1/books"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tmpl/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	books.SetupRouter(r)
	log.Info("Server started")
	log.Fatalln(r.Run())
}
