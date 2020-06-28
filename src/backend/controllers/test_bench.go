package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

func AssertExpected(t *testing.T, expected, received interface{}) {
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Got %v, wanted %v", received, expected)
	}
}

func SetupContext(db *gorm.DB) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("db", db)

	return w, c
}

func SetupRequestBody(c *gin.Context, payload interface{}) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(payload)

	c.Request = &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes.Bytes())),
	}
}
