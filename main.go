package main

import (
	_ "WebAssembly/db"
	"WebAssembly/service"
	"fmt"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/update_node", handler.AddNode).Methods("POST")
	r.HandleFunc("/add_comment", handler.AddNode).Methods("POST")
	r.HandleFunc("/get_node", handler.GetNodeById).Methods("POST")
	r.HandleFunc("/get_all_node", handler.GetAllNode).Methods("POST")
	listenAddr := getHostAndPort(address, port)
	log.Printf("Server started listening on %v", listenAddr)
	http.ListenAndServe(listenAddr, r)
}

func getHostAndPort(addr string, port int) string {
	return fmt.Sprintf("%v:%v", addr, port)
}
