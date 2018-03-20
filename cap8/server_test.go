package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handler)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("response code is %v", w.Code)
	}

	var post Post
	json.Unmarshal(w.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("cannot retrieve JSON post")
	}
}
