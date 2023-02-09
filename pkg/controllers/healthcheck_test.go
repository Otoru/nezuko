package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/otoru/nezuko/pkg/services/healthcheck"
	"go.uber.org/zap"
)

func TestHealthCheck(t *testing.T) {
	t.Run("We have a successful response by default", func(t *testing.T) {
		router := gin.Default()
		logger := zap.NewExample()
		service := healthcheck.New()

		HealthCheck(router, service, logger)

		request, err := http.NewRequest("GET", "/healthz", nil)

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		t.Logf("Expected status code: %v", http.StatusNoContent)
		t.Logf("Recived status code: %v", recorder.Code)

		if recorder.Code != http.StatusNoContent {
			t.Errorf("The received status code is not the expected one")
		}
	})

	t.Run("A failed runner generates the error status code", func(t *testing.T) {
		router := gin.Default()
		logger := zap.NewExample()
		service := healthcheck.New()

		service.Append(
			healthcheck.Runner{
				Name: "mock",
				Evaluation: func() error {
					return errors.New("mock-error")
				},
			},
		)

		HealthCheck(router, service, logger)

		request, err := http.NewRequest("GET", "/healthz", nil)

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		t.Logf("Expected status code: %v", http.StatusInternalServerError)
		t.Logf("Recived status code: %v", recorder.Code)

		if recorder.Code != http.StatusInternalServerError {
			t.Errorf("The received status code is not the expected one")
		}
	})
}
