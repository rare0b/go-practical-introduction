package pkg

import (
	"fmt"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	fmt.Fprintln(w, c)
}
