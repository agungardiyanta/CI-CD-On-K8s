package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

// GetIP extracts the user's IP address from the request
func getIPAddress(r *http.Request) string {
	// Check if the IP is coming from a reverse proxy or load balancer
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		// Check the standard forwarded header
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		// If no headers, use the remote address
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	// Handle multiple IPs in X-Forwarded-For header
	if strings.Contains(ip, ",") {
		// X-Forwarded-For can contain multiple IPs, the first one is the client IP
		ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	}
	return ip
}

// Handle the request and return the user's IP in HTML format
func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIPAddress(r)

	// Load and parse the HTML template
	tmpl, err := template.ParseFiles(filepath.Join("templates", "output-ip.html"))
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with the IP address
	data := struct {
		IP string
	}{
		IP: ip,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", ipHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
