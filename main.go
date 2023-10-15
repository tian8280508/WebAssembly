package main

import (
	_ "WebAssembly/db"
	"WebAssembly/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const (
	address = "0.0.0.0"
	port    = 8080
)

func main() {
	handler := &service.SimpleHandler{}
	r := mux.NewRouter()
	r.HandleFunc("/add_node", handler.AddNode).Methods("POST")
	r.HandleFunc("/add_edge", handler.AddEdge).Methods("POST")
	r.HandleFunc("/update_node", handler.UpdateNode).Methods("POST")
	r.HandleFunc("/add_comment", handler.AddComment).Methods("POST")
	r.HandleFunc("/get_node_version_list", handler.GetNodeVersionList).Methods("POST")
	r.HandleFunc("/get_version", handler.GetNodeVersion).Methods("POST")
	r.HandleFunc("/get_graph", handler.GetGraph).Methods("POST")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	})
	handlerWithCORS := c.Handler(r)

	listenAddr := getHostAndPort(address, port)
	log.Printf("Server started listening on %v", listenAddr)
	http.ListenAndServe(listenAddr, handlerWithCORS)
}

func getHostAndPort(addr string, port int) string {
	return fmt.Sprintf("%v:%v", addr, port)
}
