package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type Post struct {
	User    string
	Threads []string
}

//Base64URL 编码，以此来满足响应首部对 cookie值的 URL 编码要求
func showMessage(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("flash") //字符串
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no message")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello msg")
	c1 := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c1)
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler function called -" + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/set", log(setMessage))
	http.HandleFunc("/show", log(showMessage))
	server.ListenAndServe()
}
