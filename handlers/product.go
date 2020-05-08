package handlers

import (
	"log"
	"net/http"

	"github.com/unsmoker/micsroservices/data"
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
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable Read to json", http.StatusInternalServerError)
		return
	}
}
