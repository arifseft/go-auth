package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arifseft/go-auth/src/apps"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDecodeMap(t *testing.T) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/user/1", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
