package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Korisss/concet-backend/internal/domain"
	repository_mock "github.com/Korisss/concet-backend/internal/repository/mock"
	"github.com/Korisss/concet-backend/internal/service"
	"github.com/Korisss/concet-backend/internal/transport/handler"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func TestSignIn(t *testing.T) {
	// Init
	db := repository_mock.NewDBMock()
	repo := repository_mock.NewRepositoryMock(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	// Init server
	server := httptest.NewServer(handlers.InitRoutes())
	defer server.Close()

	// Init client
	client := &http.Client{}

	// Test valid request
	{
		// Create Request
		body, _ := json.Marshal(loginRequest{
			Email:    "test@example.com",
			Password: "password",
		})

		req, _ := http.NewRequest("POST", server.URL+"/auth/sign-in", bytes.NewReader(body))
		defer req.Body.Close()

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err.Error())
		}

		defer resp.Body.Close()

		// Check response
		if resp.StatusCode != http.StatusOK {
			t.Fatal("error on valid request")
		}
	}

	// Test invalid request
	{
		// Create Request
		req, _ := http.NewRequest("POST", server.URL+"/auth/sign-in", strings.NewReader(""))
		defer req.Body.Close()

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err.Error())
		}

		defer resp.Body.Close()

		// Check response
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatal("error on invalid request")
		}
	}
}

func TestSignUp(t *testing.T) {
	// Init
	db := repository_mock.NewDBMock()
	repo := repository_mock.NewRepositoryMock(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	// Init server
	server := httptest.NewServer(handlers.InitRoutes())
	defer server.Close()

	// Init client
	client := &http.Client{}

	// Test valid request
	{
		// Create Request
		body, _ := json.Marshal(domain.User{
			Name:     "name",
			Email:    "test@example.com",
			Password: "password",
		})

		req, _ := http.NewRequest("POST", server.URL+"/auth/sign-up", bytes.NewReader(body))
		defer req.Body.Close()

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err.Error())
		}

		defer resp.Body.Close()

		// Check response
		if resp.StatusCode != http.StatusOK {
			t.Fatal("error on valid request")
		}
	}

	// Test invalid request
	{
		// Create Request
		body, _ := json.Marshal(domain.User{
			Name:     "name",
			Email:    "brokenemail",
			Password: "password",
		})

		req, _ := http.NewRequest("POST", server.URL+"/auth/sign-up", bytes.NewReader(body))
		defer req.Body.Close()

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err.Error())
		}

		defer resp.Body.Close()

		// Check response
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatal("error on invalid request")
		}
	}
}
