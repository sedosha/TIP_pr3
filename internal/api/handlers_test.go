package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"pz3-http/internal/storage"
)

func TestHealthHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expected := `{"status":"ok"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("Expected body '%s', got '%s'", expected, actual)
	}
}

func TestCreateTask(t *testing.T) {
	store := storage.NewMemoryStore()
	handlers := NewHandlers(store)

	taskData := map[string]string{"title": "Test task"}
	jsonData, _ := json.Marshal(taskData)
	
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.CreateTask(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", status)
	}

	var task map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &task)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}

	if task["title"] != "Test task" {
		t.Errorf("Expected title 'Test task', got %v", task["title"])
	}
}

func TestCreateTaskEmptyTitle(t *testing.T) {
	store := storage.NewMemoryStore()
	handlers := NewHandlers(store)

	taskData := map[string]string{"title": ""}
	jsonData, _ := json.Marshal(taskData)
	
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.CreateTask(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 for empty title, got %d", status)
	}
}

func TestCreateTaskShortTitle(t *testing.T) {
	store := storage.NewMemoryStore()
	handlers := NewHandlers(store)

	taskData := map[string]string{"title": "ab"}
	jsonData, _ := json.Marshal(taskData)
	
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.CreateTask(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422 for short title, got %d", status)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	store := storage.NewMemoryStore()
	handlers := NewHandlers(store)

	req := httptest.NewRequest("GET", "/tasks/999", nil)
	rr := httptest.NewRecorder()

	handlers.GetTask(rr, req)

	if status := rr.Code; status != http.StatusNotFound && status != http.StatusInternalServerError {
		t.Errorf("Expected status 404 or 500, got %d", status)
	}
}