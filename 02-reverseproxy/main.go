package main

import (
	"io"
	"net/http"
	"net/http/httputil"

	"golang.org/x/net/websocket"
)

// ref: https://qiita.com/mackee_w/items/b6036d132888caaf5f98

func main() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:12345"
	}

	proxy := &httputil.ReverseProxy{
		Director: director,
	}
	http.Handle("/echo", proxy)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// EchoHandler ...
func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
