package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	expectedResponse := `{"email":"user@email.com","display_name":"User"}`

	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("0")

	u := newUserHandlerTest()

	if assert.NoError(t, u.getUser(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, strings.Trim(rec.Body.String(), "\n"))
	}
}

type mockUserRepository struct {
	users map[string]*db.User
}

func (u *mockUserRepository) GetUser(id string) (*db.User, error) {
	elem, ok := u.users[id]
	if ok {
		return elem, nil
	}
	return nil, &db.ErrUserNotFound{Id: id}
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: map[string]*db.User{
			"0": {
				Email: "user@email.com",
				Name:  "User",
				Id:    "0",
			},
		},
	}
}

func newUserHandlerTest() *userHandler {
	return &userHandler{
		userDB: newMockUserRepository(),
	}
}
