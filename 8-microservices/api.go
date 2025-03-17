package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{svc: svc}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/fact", s.handleGetCatFact)
	log.Println("Server started at", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
