package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestHealthCheckHandler tests the HealthCheckHandler function.
func TestHealthCheckHandler(t *testing.T) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new gin router
	r := gin.Default()

	// Register the handler
	r.GET("/health", HealthCheckHandler)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Couldn't parse response: %v\n", err)
	}

	// Check the response fields
	assert.Equal(t, "UP", response["status"])
	assert.Equal(t, "Server is running smoothly", response["message"])

	// Check the time format
	_, err = time.Parse(time.RFC3339, response["time"].(string))
	assert.NoError(t, err)
}
