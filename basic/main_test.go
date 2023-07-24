package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO https://circleci.com/blog/gin-gonic-testing/#installing-the-projects-dependencies

func SetUpRouter() *gin.Engine {
	return Router()
}

func TestPing(t *testing.T) {
	mockResponse := "pong"
	r := SetUpRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(response))
	assert.Equal(t, http.StatusOK, w.Code)

}
