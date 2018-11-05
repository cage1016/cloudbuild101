package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-zoo/bone"
	"github.com/stretchr/testify/assert"
)

func Test_Ping_Handler(t *testing.T) {
	mux := bone.New()
	mux.Get("/ping", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong"))
	}))

	r, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, "pong", string(w.Body.Bytes()))
}
