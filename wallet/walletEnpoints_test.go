package wallet

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"oblique/db"
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

	expected := `[{"title":"test mongo add wallet 2","colors":["#fff","#0000"],"_id":"5f44f87506dd2615e421a7c5"},{"title":"test mongo add wallet 2","colors":["#fff","#0000"],"_id":"5f44fef81fe31205184aba95"},{"title":"test mongo add wallet 3","colors":["#fff","#0000"],"_id":"5f44ff517de57b9bcbadbacf"}]`
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
