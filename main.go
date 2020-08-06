package main

import (
	"flag"
	"html/template"
	"log"
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
		t.temp.Execute(w, r)
	})
}

func main() {
	var addr = flag.String("addr",":8080","Application,Adress")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	log.Println("WebServe starting. Port : ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}
