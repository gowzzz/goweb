package main

import (
	"net/http"
	// "encoding/json"
	// "fmt"
	// "path"
	// "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("handleGet"))
	w.WriteHeader(200)
	return
}
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("handlePost"))
	w.WriteHeader(201)
	return
}
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("handlePut"))
	w.WriteHeader(201)
	return
}
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("handleDelete"))
	w.WriteHeader(204)
	return
}
