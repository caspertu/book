package books

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = SetupRouter(gin.Default())
	os.Exit(m.Run())
}

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"barcode":"987","title":"Go","author":"At"}`)
	req, _ := http.NewRequest("GET", "/ping", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestFetchAll(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books/", nil)
	router.ServeHTTP(w, req)
	expected := `{"data":[{"id":1,"barcode":"","title":"Go","author":"At"}]}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHandleGet(t *testing.T) {
	ID := 1
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books/"+strconv.Itoa(ID), nil)
	router.ServeHTTP(w, req)
	expected := `{"data":{"id":1,"barcode":"","title":"Go","author":"At"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHandlePost(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"barcode":"987","title":"Go","author":"At"}`)
	req, _ := http.NewRequest("POST", "/api/v1/books/", body)
	router.ServeHTTP(w, req)
	expected := `{"data":{"id":2,"barcode":"987","title":"Go","author":"At"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHandlePut(t *testing.T) {
	w := httptest.NewRecorder()
	ID := 1
	req, _ := http.NewRequest("PUT", "/api/v1/books/"+strconv.Itoa(ID), nil)
	router.ServeHTTP(w, req)
	expected := `{"data":{"id":1,"barcode":"","title":"GoGo","author":"At"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
	t.Log(w.Body.String())
}

func TestHandleDelete(t *testing.T) {
	w := httptest.NewRecorder()
	ID := 1
	req, _ := http.NewRequest("DELETE", "/api/v1/books/"+strconv.Itoa(ID), nil)
	router.ServeHTTP(w, req)
	expected := `{"data":{"id":1,"barcode":"","title":"GoGo","author":"At"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
