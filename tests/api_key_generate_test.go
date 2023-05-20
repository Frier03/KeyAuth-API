package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Frier03/KeyAuth-API/pkg/routes"
	"github.com/gin-gonic/gin"
)

func TestAPIKeyGenerateHandler(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()

	// Set up authentication routes
	routes.SetupAPIKeyRoutes(r)

	// Create request body
	reqBody := bytes.NewBufferString(`
		{
			"UUID": "a1f44782-6499-4d1e-948d-1ab37ad23b82-9a79bc5f585b0d7002a3400943585372"
		}
	`)

	// Create a test request to the login endpoint
	req, err := http.NewRequest("GET", "/api-key/generate", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	r.ServeHTTP(recorder, req)

	// Assert the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// Print the response body
	fmt.Println("Response Body:", recorder.Body.String())
}
