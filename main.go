package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jere-mie/whatsmyip/internal/utils"
)

func main() {
	port := utils.GetPort()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := utils.GetClientIP(r)
		log.Printf("Request received from %s", clientIP)
		fmt.Fprintf(w, "%s", clientIP)
	})

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
