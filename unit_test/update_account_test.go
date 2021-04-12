package unit

import (
	"bankinfo.com/routers"
	"bankinfo.com/service"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateAccountPassCase(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{ "name_bank" : "VuHH2"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/accounts/1/update", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/UpdateAccountPassCaseExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestUpdateAccountNameBankIsEmpty(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{ "name_bank" : ""}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/accounts/1/update", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content := "{\"messages\":\"Name Bank is required\"}"
	assert.Equal(t, 501, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}
