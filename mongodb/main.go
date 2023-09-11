package main

import (
	"log"
	"net/http"

	"github.com/AmanMittal2804/mongodb/router"
)

func main() {

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
