package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello Go world")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Default handler")
		d, err := ioutil.ReadAll(r.Body)
		fmt.Printf("The data is %s\n", d)
		if err != nil {
			http.Error(w, "oops Error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello %s", d)
	})
	http.HandleFunc("/GoodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bbye byt to C++ ")
	})
	http.ListenAndServe(":9090", nil)

}
