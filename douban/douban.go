package douban

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type DoubanBook struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Image       string `json:"image"`
	Catalog     string `json:"catalog"`
	Publisher   string `json:"publisher"`
	Price       string `json:"price"`
	Summary     string `json:"summary"`
	AuthorIntro string `json:"author_intro"`
	Pubdate     string `json:"pubdate"`
	Pages       string `json:"pages"`
	Isbn13      string `json:"isbn13"`
	Author      []string
}

var log = logrus.New()

func getUrl(url string) []byte {
	log.Println("Start GetUrl:", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `curl/7.64.1`)
	req.Header.Add("Accept", `*/*`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(url)
	}
	log.Println("StatusCode:", resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	//log.Println("data:", string(data))
	return data
}

//func GetBookByISBN(isbn string) (book *model.Book) {
//	url := "https://douban-api.zce.now.sh/v2/book/isbn/" + isbn
//	data := getUrl(url)
//	var b = DoubanBook{}
//	json.Unmarshal(data, &b)
//	book = model.New()
//
//	book.Barcode = b.Isbn13
//	book.Title = b.Title
//	var author string
//	for _, a := range b.Author {
//		author = author + a
//	}
//	book.Author = author
//	log.Println("Book:", book)
//	return book
//}

func New(isbn string) *DoubanBook {
	url := "https://douban-api.zce.now.sh/v2/book/isbn/" + isbn
	data := getUrl(url)
	var b = DoubanBook{}
	json.Unmarshal(data, &b)
	log.Println("Book:", b)
	return &b
}
