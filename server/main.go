package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aopal/mytholojam/server/gameplay"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request at root")
	w.Write([]byte(""))
}

// func webClientHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Serving web client")

// 	index, _ := os.Open("./web/index.html")

// 	io.Copy(w, index)
// }

// func logRoutes(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 	log.Printf("%s %s\n\n", route, router)
// 	return nil
// }

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/create-game/{gameID}", gameplay.CreateHandler)
	r.HandleFunc("/join-game/{gameID}", gameplay.JoinHandler)
	r.HandleFunc("/status/{gameID}/{actionCounter:[0-9]+}", gameplay.StatusHandler)
	r.HandleFunc("/take-action/{gameID}", gameplay.ActionHandler)
	r.PathPrefix("/play/").Handler(http.StripPrefix("/play/", http.FileServer(http.Dir("./web"))))

	gameplay.Init()

	var port string

	if len(os.Args) < 2 {
		fmt.Println("No port given, defaulting to 8080")
		port = "8080"
	} else {
		port = os.Args[1]
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
