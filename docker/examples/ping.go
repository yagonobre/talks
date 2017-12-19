package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// START OMIT
func Ping(w http.ResponseWriter, r *http.Request) {
	hn, _ := os.Hostname()
	fmt.Fprintf(w, "pong - %s", hn)
}

func main() {
	http.HandleFunc("/ping", Ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
