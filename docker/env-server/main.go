package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	startAt := time.Now()
	hostname := os.Getenv("HOSTNAME")
	port := "8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dur := time.Now().Sub(startAt)
		info := fmt.Sprintln("Hostname:", hostname, ", server stared for:", dur)
		fmt.Fprintln(w, info)
		fmt.Println(info)
	})

	fmt.Println("server stared on port:", port)

	http.ListenAndServe(":"+port, nil)
}
