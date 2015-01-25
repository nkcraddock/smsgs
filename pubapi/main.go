package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/nkcraddock/smessages/queue"
)

var publish chan queue.Event

const (
	rabbitUri = "amqp://guest:guest@172.17.0.24:5672"
	exchange  = "smsgs.evt"
	publisher = "qt"
)

func main() {
	port := flag.Int("p", 3000, "the port")
	flag.Parse()

	publish = queue.Publisher(rabbitUri, exchange)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + strconv.Itoa(*port))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	e := queue.GenerateRandomEvent(publisher)
	publish <- e
	fmt.Println("PUBLISH:", e)

}
