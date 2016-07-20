// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"code.psg.io/polymer-template/json"
)

func ShowBaseTemplate(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("templates/base.html")
	if err != nil {
		log.Printf("Unable to retrieve assets: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var req jsonrpc.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var sum int
	for _, addend := range req.Addends {
		sum += addend
	}

	resp := jsonrpc.AddResponse{Sum: sum}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
