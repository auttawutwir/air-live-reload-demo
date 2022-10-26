package main

import (
	"io"
	"log"
	"net/http"
)
func main() {
        http.HandleFunc("/", Live)

        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

func Live(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "live version 2")
}