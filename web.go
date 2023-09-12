package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string
	fmt.Print("enter the port no to listen on:")
	fmt.Scan(&port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello web")
	})
	log.Println("port is listening on", port, "......")
	http.ListenAndServe(":"+port, nil)
}
