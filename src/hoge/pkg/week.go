package pkg

import (
	"html/template"
	"net/http"
)

func Week(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/tmpl.html")
	if err != nil {
		panic(err)
	}
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}
