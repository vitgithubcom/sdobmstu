package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Создаем тестовый запрос
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// Вызываем обработчик
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверяем Content-Type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong Content-Type: got %v want application/json", contentType)
	}

	// Проверяем тело ответа
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("response body is not valid JSON: %v", err)
	}

	// Проверяем поле message
	message, ok := response["message"]
	if !ok {
		t.Error("response missing 'message' field")
	}
	if message != "hello, world!" {
		t.Errorf("wrong message: got %v want hello, world!", message)
	}
}

// Дополнительный тест для проверки метода POST (должен вернуть 405)
func TestHelloHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for POST: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}