package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user, "user should be nil")
	assert.NotNil(t, err, "Error should not be nil")
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invald login credentials", "status": "404", "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login  user", err.Message)
}
func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}
func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1","first_name": "Arturo","last_name": "Murillo","email": "amurillop90@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal user response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1,"first_name": "Arturo","last_name": "Murillo","email": "amurillop90@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Arturo", user.FirstName)
	assert.EqualValues(t, "Murillo", user.LastName)
	assert.EqualValues(t, "amurillop90@gmail.com", user.Email)
}
