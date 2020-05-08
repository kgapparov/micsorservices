package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Hello %s", string(b))
}
