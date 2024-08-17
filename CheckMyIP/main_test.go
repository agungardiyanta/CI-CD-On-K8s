package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test the GetIP function
func TestGetIP(t *testing.T) {
	tests := []struct {
		name        string
		headerKey   string
		headerValue string
		remoteAddr  string
		expectedIP  string
	}{
		{
			name:        "X-Forwarded-For header is set",
			headerKey:   "X-Forwarded-For",
			headerValue: "203.0.113.195",
			remoteAddr:  "203.0.113.195:12345",
			expectedIP:  "203.0.113.195",
		},
		{
			name:        "No X-Forwarded-For header, use RemoteAddr",
			headerKey:   "",
			headerValue: "",
			remoteAddr:  "203.0.113.195:12345",
			expectedIP:  "203.0.113.195",
		},
		{
			name:        "IPv6 address in RemoteAddr",
			headerKey:   "",
			headerValue: "",
			remoteAddr:  "[2001:db8::1]:12345",
			expectedIP:  "2001:db8::1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request
			req := httptest.NewRequest("GET", "http://example.com/", nil)

			// Set the remote address and headers if provided
			req.RemoteAddr = tt.remoteAddr
			if tt.headerKey != "" {
				req.Header.Set(tt.headerKey, tt.headerValue)
			}

			// Get the IP address using the GetIP function
			ip := getIPAddress(req)

			// Check if the IP address matches the expected result
			if ip != tt.expectedIP {
				t.Errorf("expected IP %s, got %s", tt.expectedIP, ip)
			}
		})
	}
}

// Test the ipHandler function
func TestIpHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/myip", nil)
	req.RemoteAddr = "203.0.113.195:12345"

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	ipHandler(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the response contains the expected IP address
	expectedIP := "203.0.113.195"
	if !strings.Contains(rr.Body.String(), expectedIP) {
		t.Errorf("handler returned unexpected body: got %v want IP address %v", rr.Body.String(), expectedIP)
	}

	// Check if the response contains the HTML structure
	if !strings.Contains(rr.Body.String(), "<div class=\"ip-container\">") {
		t.Errorf("handler returned unexpected body: missing expected HTML structure")
	}
}
