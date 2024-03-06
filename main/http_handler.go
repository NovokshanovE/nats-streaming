package main

import (
	"fmt"
	"log"
	"net/http"
)

func showOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }
	if val, ok := cache.Orders[id]; ok {
		fmt.Fprint(w, cache.to_json(id))
		fmt.Print(val)
	} else {
		err := fmt.Errorf("Error: try to get order by ID = %s", id)
		log.Panic(err)
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}
}

func runHttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/order", showOrder)

	log.Println("Server run on: http:localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
