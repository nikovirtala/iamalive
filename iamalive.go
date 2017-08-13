package main

import (
    "fmt"
    "html/template"
    "log"
    "net"
    "net/http"
    "strings"
    "os"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
    var name, _ = os.Hostname()
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello! Your request was processed by: %s", name)
}

func test(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    var status string
    if r.Method == "GET" {
        t, _ := template.ParseFiles("test.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("destination:", r.Form["destination"])
        destination := r.FormValue("destination")
        conn, err := net.Dial("tcp", destination)
        if err != nil {
                fmt.Fprintf(w, "Connection error:", err)
                status = " is unreachable."
        } else {
                status = " is online!"
                defer conn.Close()
        }
        fmt.Fprintf(w, destination + status)
    }
}

func main() {
    http.HandleFunc("/", sayhello)
    http.HandleFunc("/test", test)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
