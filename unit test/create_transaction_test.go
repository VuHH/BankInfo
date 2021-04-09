package unit

import (
	"../routers"
	"../service"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTransactionPassCase(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": 3,"amount": 112000,"transaction_type": "deposit"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransacitonJsonExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionFailedAccountID(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": "3","amount": 112000,"transaction_type": "deposit"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionFailedAccountIDExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionFailedAmount(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": 3,"amount": "112000","transaction_type": "deposit"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionFailedAmountExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionFailedTransactionType(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": 3,"amount": 112000,"transaction_type": ""}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionFailedTransactionTypeExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionWrongTransactionType(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": 3,"amount": 112000,"transaction_type": "aaaaaa"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionWrongTransactionTypeExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionAmountIsNegative(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": 3,"amount": -10000,"transaction_type": "deposit"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionAmountIsNegativeExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}

func TestCreateTransactionAccountIdIsNegative(t *testing.T) {

	token := service.GenerateToken("admin", true)
	var bearer = "Bearer " + token

	r := routers.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"account_id": -3,"amount": 10000,"transaction_type": "deposit"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users/1/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	content, err := ioutil.ReadFile("./expected_result/CreateTransactionAccountIdIsNegativeExpect.json")

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(content), w.Body.String())
}
