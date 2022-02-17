package db

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDBInitialization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := ConnectDB()
	if db == nil {
		t.Fail()
	}
}