package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	l.Println("Recieved Terminate, Graceful Shutdown", sig)
	// Gracefully ShutDown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
