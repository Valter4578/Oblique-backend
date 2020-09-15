package wallet

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

func TestGetWallets(t *testing.T) {
	password := os.Getenv("URI")
	fmt.Println(password)
	db.ConnectDB(password)

	req, err := http.NewRequest("GET", "/wallets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWallets)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"title":"test mongo add wallet 2","colors":["#fff","#0000"],"_id":"5f44f87506dd2615e421a7c5"},{"title":"test mongo add wallet 2","colors":["#fff","#0000"],"_id":"5f44fef81fe31205184aba95"},{"title":"unit test AddOperation","colors":["#fff","#ffff"],"_id":"5f60dc95a84e442866ef66ca"},{"title":"unit test AddOperation","colors":["#fff","#ffff"],"_id":"5f60de0b7ff9ff63c339c83b"},{"title":"unit test AddOperation","colors":["#fff","#ffff"],"_id":"5f60de54cb1f87b682f2efce"}]`
	body := rr.Body.String()

	require.JSONEq(t, expected, body)
}

func TestGetWallet(t *testing.T) {
	password := os.Getenv("URI")
	fmt.Println(password)
	db.ConnectDB(password)

	req, err := http.NewRequest("GET", "/wallet", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{"id": "5f44f87506dd2615e421a7c5"}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWallet)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"title":"test mongo add wallet 2","colors":["#fff","#0000"],"_id":"5f44f87506dd2615e421a7c5"}`
	body := rr.Body.String()

	require.JSONEq(t, expected, body)
}

func TestAddWallet(t *testing.T) {
	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	jsonStr := []byte(`{"title": "unit test AddWallets","colors": ["#fff","#ffff"]}`)

	req, err := http.NewRequest("POST", "/operation", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(AddWallet)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	t.Log(rr.Body.String())

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDeleteWallet(t *testing.T) {
	uri := os.Getenv("URI")
	fmt.Println(uri)
	db.ConnectDB(uri)

	req, err := http.NewRequest("DELETE", "/wallet", nil)
	if err != nil {
		t.Fatal(err)
	}

	// get the last operation's id
	var wallets []model.Wallet
	db.GetWallets(&wallets)
	id := wallets[len(wallets)-1].ID

	t.Log(id.String())
	vars := map[string]string{"id": id.String()[1 : len(id)-1]}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteWallet)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
