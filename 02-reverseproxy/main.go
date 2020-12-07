package main

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// ref: https://qiita.com/mackee_w/items/b6036d132888caaf5f98

func main() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:12345"
		req.Header.Add("X-Test-Header", time.Now().Format(time.RFC3339))
	}

	mod := func(res *http.Response) error {
		res.Header.Set("Access-Control-Allow-Origin", "*")
		res.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		return nil
	}

	proxy := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: mod,
	}

	err := http.ListenAndServe(":8080", proxy)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
