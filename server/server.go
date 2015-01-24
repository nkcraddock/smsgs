package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
)

func main() {
	port := flag.Int("p", 3000, "the port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + strconv.Itoa(*port))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {

}
