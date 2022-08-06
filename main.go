package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"microservices/handlers"
)

func main() {
	// As we are going to use our own handler so below one will be commented
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling the client request and ready for Response ")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(http.StatusBadRequest)
		}
		fmt.Fprintf(w, "hello %s", d)

	})
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	log.Println("Start the Server")
	s := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
	log.Println("Server Started ")
}
