package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionA(t *testing.T) {
	t.Log("Running TestFunctionA in file1_test.go")
	// Create a test request
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Fatalf("Running TestFunctionA in file1_test.go")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	// Assert the response
	assert.Equal(t, string(body), "{\"message\":\"Rohit\"}")
}
