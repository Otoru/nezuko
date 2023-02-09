package healthcheck

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestService(t *testing.T) {
	t.Run("We have a successful response by default", func(t *testing.T) {
		service := New()
		logger := zap.NewExample()

		result := service.Validate(logger)

		t.Logf("Recived status code: %s", result)

		if result != nil {
			t.Errorf("The received result is not the expected one")
		}
	})

	t.Run("A failed runner generates the error on validate method", func(t *testing.T) {
		service := New()
		logger := zap.NewExample()

		mock := func() error {
			return errors.New("mock-error")
		}

		service.Append("mock", mock)

		result := service.Validate(logger)

		t.Logf("Recived status code: %s", result)

		if result == nil {
			t.Errorf("The received result is not the expected one")
		}
	})
}
