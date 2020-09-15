package operation

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"oblique/db"
	"oblique/model"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestAddOperation(t *testing.T) {
	t.Log("TestAddOperation")

	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	jsonStr := []byte(`{"title": "test mongo insert operation to category 1001", "amount": 1001}`)

	req, err := http.NewRequest("POST", "/operation", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(AddOperation)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	t.Log(rr.Body.String())

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestGetOperations(t *testing.T) {
	t.Log("TestGetOperations")

	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	req, err := http.NewRequest("GET", "/operations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOperations)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteOperation(t *testing.T) {
	t.Log("TestDeleteOperation")

	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	req, err := http.NewRequest("DELETE", "/operation", nil)
	if err != nil {
		t.Fatal(err)
	}

	// get the last operation's id
	var operations []model.Operation
	db.GetOperations(&operations)
	id := operations[len(operations)-1].ID

	t.Log(id.String())
	vars := map[string]string{"id": id.String()[1 : len(id)-1]}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteOperation)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetOperation(t *testing.T) {
	t.Log("TestGetOperation")
	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	req, err := http.NewRequest("GET", `/operation/`, nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{"id": "5f54fdb89541a3b42f61ef83"}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOperation)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"title":"post test mongo operation 1","amount":123431,"time":"0001-01-01T00:00:00Z","_id":"5f54fdb89541a3b42f61ef83"}`
	body := rr.Body.String()

	require.JSONEq(t, expected, body)
}
