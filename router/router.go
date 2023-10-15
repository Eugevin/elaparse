package router

import (
	"fmt"
	"elaparse/parser"
	"net/http"
)

func Init() {
  fs := http.FileServer(http.Dir("./public"))

  http.Handle("/", fs)

  http.HandleFunc("/api", parser.HandleApi)
  http.HandleFunc("/api/close", parser.HandleApiClose)

  fmt.Print("\nWelcome to Elaparse!\n\n")
  fmt.Print("Open http://localhost:4000 in your browser window\n\n")

  http.ListenAndServe("localhost:4000", nil)

}
