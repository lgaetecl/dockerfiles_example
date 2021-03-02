package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Body struct {
	Key string
}

func TestBodyLogging(t *testing.T) {

	body := &Body{
		Key: "Value",
	}

	jsonBody, _ := json.Marshal(body)
	if req, err := http.NewRequest("Post", "/", bytes.NewBuffer(jsonBody)); err == nil {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(BodyLogging)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)

		}

	} else {
		t.Fatal(err)
	}

}
