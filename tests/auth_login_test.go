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

func TestAuthLoginHandler(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()

	// Set up authentication routes
	routes.SetupAuthRoutes(r)

	// Create request body
	reqBody := bytes.NewBufferString(`
		{
			"username": "John",
			"password": "Doe"
		}
	`)

	// Create a test request to the login endpoint
	req, err := http.NewRequest("POST", "/auth/login", reqBody)
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
