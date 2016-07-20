// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ShowBaseTemplate).Methods("GET")
	router.HandleFunc("/adder", ShowBaseTemplate).Methods("GET")
	router.HandleFunc("/add", Add).Methods("POST")

	// Prefix the bindata filesystem
	fs := assetFS()
	fs.Prefix = "/static/"
	router.NotFoundHandler = http.FileServer(fs)

	log.Fatal(http.ListenAndServe(":8080", router))
}
