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
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello! Your request was processed by: %s \n", name)
    log.Print("Served request ",r,"\n")
}

func test(w http.ResponseWriter, r *http.Request) {
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
        log.Print("Served request ",r,"\n")
    }
}

func main() {
  log.SetOutput(os.Stderr)
  log.Println("Starting server ...")
    http.HandleFunc("/", sayhello)
    http.HandleFunc("/test", test)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
