package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Devopshift app backend live and responding!!\n")
}

func main() {
	log.Print("Devopshift app server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}
