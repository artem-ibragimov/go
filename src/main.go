package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

type IProxyList map[string]string
type Config struct {
	Proxy IProxyList `json:"Proxy"`
}

func main() {
	var config Config
	data, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(data, &config)
	fmt.Println(config)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})
}
