package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func setup() {
	//mux router with added question routes
	// m = mux.NewRouter()

	// DB := database.Init()
	// h := routes.New(DB)

	// mainRouter := mux.NewRouter()
	// authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestGet400(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/questions/1/SC", nil)
	if err != nil {
		t.Fatal("Creating 'GET /questions/1/SC' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
	}
}
