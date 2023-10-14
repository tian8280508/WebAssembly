package service

import (
	"WebAssembly/db"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type SimpleHandler struct {
	ctx context.Context
}

func parsePayload(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(payload)
}

func (*SimpleHandler) AddNode(w http.ResponseWriter, r *http.Request) {
	p := &AddNodePayload{}
	if err := parsePayload(r, p); err != nil {
		http.Error(w, "AddNode Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if p.Name == "" {
		http.Error(w, "AddNode Incomplete data: name missing", http.StatusBadRequest)
		return
	}

	if err := db.CreateNode(p.Name, p.Content); err != nil {
		log.Printf("db.CreateNode error:%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("AddNode Success! parentId: %v, Name: %s", p.ParentId, p.Name)
}

func (*SimpleHandler) GetAllNode(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetNodeBatch()
	if err != nil {
		log.Printf("GetAllNodePayload error:%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (*SimpleHandler) GetNodeById(w http.ResponseWriter, r *http.Request) {
	p := &GetNodeByIdPayload{}
	if err := parsePayload(r, p); err != nil {
		http.Error(w, "GetNodeByIdPayload Failed to decode JSON", http.StatusBadRequest)
		return
	}

	res, err := db.GetOneNode(p.Id)
	if err != nil {
		log.Printf("db.GetOneNode error:%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (*SimpleHandler) UpdateNode(w http.ResponseWriter, r *http.Request) {
	p := &UpdateNodePayload{}
	if err := parsePayload(r, p); err != nil {
		http.Error(w, "UpdateNode Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := db.UpdateNode(p.Id, p.Name, p.Content); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (*SimpleHandler) AddComment(w http.ResponseWriter, r *http.Request) {

}

func (*SimpleHandler) GetNodeVersionList(w http.ResponseWriter, r *http.Request) {
	p := &GetNodeVersionListPayload{}
	if err := parsePayload(r, p); err != nil {
		http.Error(w, "GetNodeVersionList Failed to decode JSON", http.StatusBadRequest)
		return
	}

	res, err := db.GetNodeVersionList(p.NodeId)
	if err != nil {
		log.Printf("db.GetOneNode error:%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (*SimpleHandler) AddEdge(w http.ResponseWriter, r *http.Request) {
	p := &AddEdgePayload{}
	if err := parsePayload(r, p); err != nil {
		http.Error(w, "AddEdge Failed to decode JSON", http.StatusBadRequest)
		return
	}
	if err := db.CreateEdge(p.SrcId, p.TarId, p.Desc); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
