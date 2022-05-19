package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var socketDir = "/socket"

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page, Websocket at "+socketDir)
}

func main() {

	// Setup signal monitor
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	http.HandleFunc(socketDir, wsHandler)
	http.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
