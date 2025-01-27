package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient wraps the standard http.Client to simplify API requests.
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates and returns a new HTTPClient with default settings.
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second, // Timeout for requests
		},
	}
}

// Get sends a GET request to the specified URL and returns the response body or an error.
func (h *HTTPClient) Get(url string, headers map[string]string) ([]byte, error) {
	return h.doRequest(http.MethodGet, url, nil, headers)
}

// Post sends a POST request to the specified URL with the provided payload and returns the response body or an error.
func (h *HTTPClient) Post(url string, payload interface{}, headers map[string]string) ([]byte, error) {
	return h.doRequest(http.MethodPost, url, payload, headers)
}

// Put sends a PUT request to the specified URL with the provided payload and returns the response body or an error.
func (h *HTTPClient) Put(url string, payload interface{}, headers map[string]string) ([]byte, error) {
	return h.doRequest(http.MethodPut, url, payload, headers)
}

// Delete sends a DELETE request to the specified URL and returns the response body or an error.
func (h *HTTPClient) Delete(url string, headers map[string]string) ([]byte, error) {
	return h.doRequest(http.MethodDelete, url, nil, headers)
}

// doRequest is a helper function for making HTTP requests.
func (h *HTTPClient) doRequest(method, url string, payload interface{}, headers map[string]string) ([]byte, error) {
	var body io.Reader

	// Encode payload if provided
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %v", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	// Create a new request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set default Content-Type for payload requests
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send the request
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("non-2xx status code: %d, response: %s", resp.StatusCode, responseBody)
	}

	return responseBody, nil
}
