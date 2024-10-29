package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandleHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hay")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HandleHttp(recorder, request)

	// mengecek hasil test
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}