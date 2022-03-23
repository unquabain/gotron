package main

import (
  "net/http"
  "flag"
  "fmt"
  "os"
)

var fsRoot string
var port int

func Init() {
  flag.StringVar(&fsRoot, `root`, `.`, `the root of the file system to serve`)
  flag.IntVar(&port, `port`, 6066, `the port to serve on`)
  flag.Parse()
  http.Handle(`/`, http.FileServer(http.Dir(fsRoot)))
}

func main() {
  Init()
  fmt.Println(`serving on :`, port)
  err := http.ListenAndServe(fmt.Sprintf(`:%d`, port), nil)
  if err != nil {
    fmt.Fprintf(os.Stderr, `failed to start server: %v`, err)
  }
}
