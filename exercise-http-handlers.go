package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  // your http.Handle calls here
  http.Handle("/string", String("I'm a frayed knot."))
  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
  log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s)
}

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, fmt.Sprintf("%v%v%v", s.Greeting, s.Punct, s.Who))
}
