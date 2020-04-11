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
	r.Static("/static", "./static")  // 静态文件
	r.LoadHTMLGlob("views/*") // 模版目录
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	r.GET("/", showIndex)
	books.SetupRouter(r)
	log.Info("Server started")
	log.Fatalln(r.Run())
}

type Teacher struct {
	Name    string
	Subject string
}
type Student struct {
	Id      int
	Name    string
	Country string
}
type Rooster struct {
	Teacher  Teacher
	Students []Student
}

func showIndex(c *gin.Context) {
	teacher := Teacher{
		Name:    "Alex",
		Subject: "Physics",
	}
	students := []Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	rooster := Rooster{
		Teacher:  teacher,
		Students: students,
	}
	c.HTML(http.StatusOK, "layout.gohtml", rooster)
}