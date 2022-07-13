package main

import (
    "fmt"
    "net/http"
    "os"
)

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Payments microservice version", os.Getenv("SERVICE_VERSION"))
}

func main() {
	fmt.Println("Starting web-payments version", os.Getenv("SERVICE_VERSION"))
    http.HandleFunc("/", paymentsHandler)
    http.ListenAndServe(":8080", nil)
}
