package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	temp     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.temp =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		t.temp.Execute(w, nil)
	})
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8080", nil);
		err != nil {
		http.ListenAndServe("ListenAndServe:", nil)
	}
}
