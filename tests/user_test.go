package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func Get(uri string, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestFetchAllUser(t *testing.T) {
	var w *httptest.ResponseRecorder
	urlIndex := "/user"
	w = Get(urlIndex, router)
	fmt.Println(w.Code, w.Body.String())
}
