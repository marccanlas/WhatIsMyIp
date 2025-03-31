package utils

import (
	"net/http"
	"os"
	"strings"
)

func GetPort() string {
	port := os.Getenv("WHATSMYIP_PORT")
	if port == "" {
		// Default to port 8080 if the environment variable is not set
		port = "8080"
	}
	return port
}

func GetClientIP(r *http.Request) string {
	// Extract the client's IP address from the request
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
