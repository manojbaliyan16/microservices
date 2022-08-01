package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		// Ready the Body from request
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//	w.WriteHeader(http.StatusBadRequest)
			//	w.Write([]byte("oops"))
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello %s", d)
		//log.Printf("Data %s\n", d)
	})
	http.HandleFunc("/GoodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye  world")
	})
	http.ListenAndServe(":9090", nil) // DEFAULT HANDLER

}
