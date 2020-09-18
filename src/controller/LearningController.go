package controller

import (
	"log"
	"net/http"
)

func HelloController() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
