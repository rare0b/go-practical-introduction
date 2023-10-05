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
	http.HandleFunc("/set-cookie", pkg.SetCookie)
	http.HandleFunc("/get-cookie", pkg.GetCookie)
	http.HandleFunc("/week", pkg.Week)
	http.HandleFunc("/map-week", pkg.MapWeek)
	server.ListenAndServe()
}
