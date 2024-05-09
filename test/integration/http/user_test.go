package http

import (
	"net/http"
	"testing"
)

func TestGetUsers(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/v1/users")
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}

func TestGetUserByID(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/v1/users/1")
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}
