package main

import "net/http"

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type aboutHandler struct{}

func (ah *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func file() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "folder")
	// })
	// http.ListenAndServe(":8080", nil) // DefaultServerMux")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("floder")))
}

func main() {
	mh := &myHandler{}
	ah := &aboutHandler{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // DefaultServeMux
		// Handler: mh,
	}

	http.Handle("/hello", mh)
	http.Handle("/about", ah)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe() // 与下面一行等价
	// http.ListenAndServe("localhost:8080", nil) // DefaultServerMux

	// http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
}
