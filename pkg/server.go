package main

import (
	"fmt"
	"net/http"
	"strings"
)

//Server ...
type Server struct {
	url string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/edit/"):]
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte("it works"))
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func Start(s Server) {
	http.HandleFunc("/view/", handler)
	http.HandleFunc("/edit/", handler)
	http.HandleFunc("/save/", handler)
	fmt.Println(http.ListenAndServe(s.url, nil))
}

func main() {
	s := Server{url: "https://g.codefresh.io"}
	t := [6]int{2, 3, 5, 7, 11, 13}
	tt := t[0:4]
	fmt.Println(tt)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println(s)
}
