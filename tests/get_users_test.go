package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arifseft/go-auth/src/utils/response"

	"github.com/arifseft/go-auth/src/apps"
	db "github.com/arifseft/go-auth/src/database"
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/arifseft/go-auth/src/repositories"
	"github.com/arifseft/go-auth/src/utils/flag"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func initTestGetUsers() (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	userRepository := repositories.UserRepository{Conn: db.GetDB().Table("users")}
	userRepository.CreateUser(entity.User{
		ID:       1,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email1@email.com",
		Name:     "M. Arif Sefrianto",
		Password: "123456",
	})
	userRepository.CreateUser(entity.User{
		ID:       2,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email2@email.com",
		Name:     "M. Arif Sefrianto",
		Password: "123456",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	return w, r
}

func TestGetUsers(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestGetUsers()
		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.Equal(t, flag.GetUsersSuccess.Message, actual.Message)
		assert.NotEmpty(t, actual.Data)
		assert.Empty(t, actual.Errors)
	})
}
