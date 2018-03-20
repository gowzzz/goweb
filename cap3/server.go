package main

import (
	"fmt"
	"net/http"
)

/*
https://127.0.0.1:8080/das    =>    hello:das
*/
type Myhandler1 struct{}

func (m *Myhandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Myhandler1:%s", r.URL.Path[1:])
}

type Myhandler2 struct{}

func (m *Myhandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Myhandler2:%s", r.URL.Path[1:])
}
func main() {
	my1 := Myhandler1{}
	my2 := Myhandler2{}
	server := &http.Server{
		Addr: ":8080",
		// Handler: &my, 默认使用默认的多路复用器
	}
	http.Handle("/my1", &my1)
	http.Handle("/my2", &my2)

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
