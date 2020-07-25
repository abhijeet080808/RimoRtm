// Package webserver provides a WebServer
package webserver

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	"net/http"

	// go get github.com/gorilla/websocket
	"github.com/gorilla/websocket"
)

// WebServer serves the current status to all connected web clients
// using web sockets.
type WebServer struct {
	host       string
	port       int
	httpServer *http.Server
	upgrader   *websocket.Upgrader
}

// New creates a new WebServer on specified host and port
func New(host string, port int) *WebServer {

	server := WebServer{
		host: host,
		port: port,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/websocket", server.HandleWebSocketRequest)
	mux.HandleFunc("/", server.HandleIndexRequest)

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}

	// Start on a separate goroutine
	go server.Start()

	return &server
}

// HandleIndexRequest handles HTTP request for index page
func (s *WebServer) HandleIndexRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "html/index.html")
}

// HandleWebSocketRequest handles HTTP request for web socket
func (s *WebServer) HandleWebSocketRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Client subscribed")

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}

	num := 0

	for {
		time.Sleep(2 * time.Second)

		num += 2

		numJSON, err := json.Marshal(num)
		if err != nil {
			log.Panicln(err)
		}

		err = conn.WriteMessage(websocket.TextMessage, numJSON)
		if err != nil {
			log.Println(err)
			break
		}
	}

	conn.Close()
	log.Println("Client unsubscribed")
}

// Start will start the web server
func (s *WebServer) Start() {
	log.Println("Starting HTTP server at", s.host, s.port)
	// Blocking call
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
