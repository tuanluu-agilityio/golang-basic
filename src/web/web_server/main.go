package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm() // parse arguments
  fmt.Println(r.Form) // print Form information in server side
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("value: ", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello World!!!") // send data to client side
}

func main() {
  http.HandleFunc("/", sayHelloName) // set router
  err := http.ListenAndServe(":9090", nil) // set listen port
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
