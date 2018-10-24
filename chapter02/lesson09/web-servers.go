package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	Msg string
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", h.Msg, " ,Header:", r.Header)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "String, ", s)
}

func (st Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Struct, ", st)
}

// Web 服务器
func main() {

	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))

}
