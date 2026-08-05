package main

import (
	"fmt"
	"net/http"
	"time"
)

func printUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User123")
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, world") })
	http.Handle("GET /user", Middleware(http.HandlerFunc(printUser)))	
	fmt.Println("Сервер запущен на 5000 порту")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ping := time.Now()

		next.ServeHTTP(w, r)

		ping2 := time.Since(ping)
		fmt.Printf("Ping %s\n",ping2)

	})
}
