package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nishilfaldu/cloudproof/internal/api"
)

func main() {
	http.HandleFunc("/api/findings", api.FindingsHandler)
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
