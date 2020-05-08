package handlers

import (
	"log"
	"net/http"
)

//Pruducts ...
type Pruducts struct {
	l *log.Logger
}

//NewProduct ...
func NewProduct(l *log.Logger) *Pruducts {
	return &Pruducts{l}
}

//ServeHTTP ...
func (p *Pruducts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.NewProduct()
}
