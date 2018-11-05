package main // import "github.com/cage1016/bonego"

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()

	mux.Get("/", http.HandlerFunc(RootHandler))
	mux.Get("/ping", http.HandlerFunc(PingHandler))

	http.ListenAndServe(":8080", mux)
}

func PingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("pong"))
}

func RootHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("helle world"))
}
