package userservice

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"bytes"

	"genz_server/config/db"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"email": "test@gmail.com",
		"password": "Test@123",
		"is_email_verified": false,
		"social_profile": "{}",
	})

	if err != nil {
		t.Fail()
	}

	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPost, "api/v1/users/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	addUser := AddUserHandler(database)
	addUser(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	expected, actual := 201, w.Code
	if expected == actual {
		t.Logf("Create User Success")
	} else {
		t.Fail()
	}
}

func TestLoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"email": "test@gmail.com",
		"password": "Test@123",
	})

	if err != nil {
		t.Fail()
	}

	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPost, "api/v1/users/login", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	loginUser := UserLoginHandler(database)
	loginUser(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	expected, actual := 200, w.Code
	if expected == actual {
		t.Logf("user logged in")
	} else {
		t.Errorf("Expected code %q but got %q", expected, actual)
	}

}

func TestGetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key: "id",
			Value: "620e7780665f45729159d720", // change the user id you want
		},
	}
	getUserById := GetUserByIdHandler(database)
	getUserById(c)

	expected, actual := 302, w.Code // change the code according to the user id
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}
}