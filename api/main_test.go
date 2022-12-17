package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
