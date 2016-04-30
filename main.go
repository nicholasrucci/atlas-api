package main

import (
	"net/http"
	"log"
)



func main() {
	log.Fatal(http.ListenAndServe(":3000", nil))
}
