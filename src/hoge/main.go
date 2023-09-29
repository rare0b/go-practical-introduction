package main

import (
	"hoge/pkg"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", pkg.Body)
	http.HandleFunc("/set_cookie", pkg.SetCookie)
	http.HandleFunc("/get_cookie", pkg.GetCookie)
	server.ListenAndServe()
}
