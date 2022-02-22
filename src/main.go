package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("https://rs.aspsp.ob.forgerock.financial:443")
	log.Printf("forwarding to -> %s%s\n", target.Scheme, target.Host)

	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(w, req)
	})

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}
