package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func main() {

	log.Println("Starting home-site server")
	const port = 8081

	//frontendui := ""

	ui := flag.String("ui", "", "ui location")
	flag.Parse()
	if *ui != "" {
		log.Printf("Address: http://localhost:%v UI dir is %v", port, *ui)
	} else {
		log.Fatalln("No -ui flag. Please supply the location of the frontend. e.g -ui=\"../frontend/dist\"")
	}
	r := httprouter.New()
	//	r.GET("/", index)
	r.Handler("GET", "/", http.FileServer(http.Dir(*ui)))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "Working")
}
