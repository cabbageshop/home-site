package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {

	log.Println("Starting home-site server")
	const port = 8081

	ui := flag.String("ui", "", "ui location")
	flag.Parse()
	if *ui != "" {
		http.Handle("/", http.FileServer(http.Dir(*ui)))
		log.Printf("Address: http://localhost:%v UI dir is %v", port, *ui)
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	} else {
		log.Fatalln("No -ui flag. Please supply the location of the frontend. e.g -ui=\"../frontend/dist\"")
	}
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Working")
}
