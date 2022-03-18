package main

import (
	"integral_test/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/joke", handler.GetJoke)
	mux.HandleFunc("/joke", handler.Index)

	http.ListenAndServe(":1212", mux)

}
