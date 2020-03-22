package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllNote(t *testing.T) {
	req, err := http.NewRequest("GET", "/note", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
	rr := httptest.NewRecorder()

	noteHandler := NoteHandler{
		DB:     db,
		Engine: r,
	}

	handler := http.HandlerFunc(GetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"title":"exercise","completed":"false","category_id":"0"]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
