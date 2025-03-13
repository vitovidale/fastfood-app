package logger_test

import (
	"testing"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/config"
	"github.com/vitovidale/fastfood-app/internal/adapter/logger"
)

func TestSetLogger(t *testing.T) {
	mockConfig := &config.App{
		// Add necessary fields for the mock configuration
	}

	// Call the Set function
	logger.Set(mockConfig)

	// Verify the logger is initialized correctly
	// This might involve checking a global variable or other side effects
	if logger.GetLoggerInstance() == nil {
		t.Error("Expected logger to be initialized, but it was nil")
	}
}
