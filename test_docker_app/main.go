package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Docker App\nVersion: %s\nOS: %s\nArch: %s\n",
			os.Getenv("APP_VERSION"),
			runtime.GOOS,
			runtime.GOARCH,
		)
	})
	port := "0.0.0.0:8080"
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(port, nil)
}
