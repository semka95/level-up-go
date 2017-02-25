package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})

	mux.HandleFunc("/articles/latest/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/latest/")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go Server")
		fmt.Fprintf(w, `
		<html>
			<body>
				Hello Gopher
			</body>
		</html>`)
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})

	mux.HandleFunc("/error/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error", 500)
	})

	http.ListenAndServe(":3000", mux)
}
