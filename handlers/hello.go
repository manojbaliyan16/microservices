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

// idematic Principle to define code

// dependency Injection in go for for logger
// the benifits of the dependency injection is this we can replace l --> *log.Logger woth any orher value
// helps us to wrote fast unit testing
func NewHello(l *log.Logger) *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Default handler") // try to avoid directly concrete object inside the hadnler so needs to change the looger
	d, err := ioutil.ReadAll(r.Body)
	h.l.Printf("The data is %s\n", d)
	if err != nil {
		http.Error(w, "oops Error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello %s", d)

}
