package unit

import (
	"../routers"
	"../service"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginWithNoUsernameAndPassword(t *testing.T) {

	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"username" : "","password" : ""}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	content := "{\"error\":\"Key: 'Login.Username' Error:Field validation for 'Username' failed on the 'required' tag\\nKey: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag\"}"

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, content, w.Body.String())

}

func TestLoginWithNoUserName(t *testing.T) {
	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"username" : "","password" : "123"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	content := "{\"error\":\"Key: 'Login.Username' Error:Field validation for 'Username' failed on the 'required' tag\"}"

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, content, w.Body.String())
}

func TestLoginWithNoPassword(t *testing.T) {
	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"username" : "123","password" : ""}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	content := "{\"error\":\"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag\"}"

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, content, w.Body.String())
}

func TestLoginWithAccountNotExist(t *testing.T) {
	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"username" : "vuhh1","password" : "123"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	content := "{\"message\":\"Account not exist\"}"

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, content, w.Body.String())
}

func TestLoginWithAccountExist(t *testing.T) {
	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"username" : "admin","password" : "admin"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	content := "{\"token\":\"" + service.GenerateToken("admin", true) + "\"}"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, content, w.Body.String())
}
