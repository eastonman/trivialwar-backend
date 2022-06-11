package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	flag "github.com/spf13/pflag"
)

var upgrader = websocket.Upgrader{}

var socketDir = "/socket"

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page, Websocket at "+socketDir)
}

var ServerAddr *string = nil

func init() {
	// command line args definition
	ServerAddr = flag.String("listen", ":8080", "The listen address for the server")

}

func main() {

	// Parse command line args
	flag.Parse()

	// Setup signal monitor
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	// Router
	router := mux.NewRouter()
	router.HandleFunc(socketDir, wsHandler)
	router.HandleFunc("/", homePageHandler)

	srv := &http.Server{
		Addr:    *ServerAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-interruptChan

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}
