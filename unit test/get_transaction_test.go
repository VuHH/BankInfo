package unit_test

import (
	"../routers"
	"../service"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpStatus int
var message interface{}

func TestGetTransactionsService(t *testing.T) {

	// User Id is empty, Account not empty
	httpStatus, message = service.GetTransactions("", "1")
	if httpStatus != 200 {
		t.Error(message)
	}

	// User Id is empty, Account empty
	httpStatus, message = service.GetTransactions("", "")
	if httpStatus != 200 {
		t.Error(message)
	}
}

func TestGetTransactionWithNoAccountId(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/1/transactions/", nil)
	req.Header.Add("Authorization", bearer)
	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/GetTransactionJsonExpect.json")

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(content), w.Body.String())

}

func TestGetTransactionWithAccountId(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/1/transactions/?account_id=1", nil)
	req.Header.Add("Authorization", bearer)
	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/GetTransactionJsonExpectWithAccountId.json")

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(content), w.Body.String())

}

func TestGetTransactionWithNoExist(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/1/transactions/?account_id=4", nil)
	req.Header.Add("Authorization", bearer)
	r.ServeHTTP(w, req)

	content := "{\"messages\":\"No Records\"}"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, content, w.Body.String())

}
