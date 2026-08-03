/* 
 * File: backend/pkg/api/handlers.go
 * Author: Tariq Scott
 * Date: 2026-08-02
 * Description: Defines the HTTP handlers for the API
 * References:
 * - https://grafana.com/blog/how-i-write-http-services-in-go-after-13-years/
 */

 package api

 import (
	"encoding/json"
	"net/http"

	"github.com/tariqsco/kubegazer/backend/pkg/k8s"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // throw a 500 error
		return
	}
}
 
func (s *Server) handleListNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := k8s.ListNodes(r.Context(), s.clientset)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, nodes)
}