package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{}
	payload := map[string]any{
		"title":       "Learn DDD",
		"description": "Get hands-on with Domain-Driven Design",
		"priority":    8000,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/tasks",
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.handleCreateTask)
	handler.ServeHTTP(rec, req)
	t.Logf("Response status: %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var resBody map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if resBody["title"] != payload["title"] {
		t.Errorf("Expected title %s, got %s", payload["title"], resBody["title"])
	}
}
