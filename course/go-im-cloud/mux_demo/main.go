package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("hello, %s\n", r.URL.Query().Get("user"))
		fmt.Print("write:", msg)
		w.Write([]byte(msg))
		time.Sleep(time.Second * 3)
		msg = "my name is vincent\n"
		fmt.Print("write:", msg)
		w.Write([]byte(msg))
		msg = "my name is king\n"
		fmt.Print("write:", msg)
		w.Write([]byte(msg))
		msg = "my name is ralf\n"
		fmt.Print("write:", msg)
		w.Write([]byte(msg))
	})
	http.ListenAndServe(":8000", mux)
}
