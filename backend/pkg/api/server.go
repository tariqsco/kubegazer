/*
 * File:
 * Author: Tariq Scott
 * Date: 2026-07-30
 * Description: API Server lifecycle, route registration,
 *              and graceful shutdown for KubeGazer.
 */

package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"k8s.io/client-go/kubernetes"
)

type Server struct {
	clientset kubernetes.Interface
	mux       *http.ServeMux
}

func NewServer(clientset kubernetes.Interface) *Server {
	s := &Server{
		clientset: clientset,
		mux:       http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /api/v1/nodes", s.handleListNodes)
}

// Handler stub for listing nodes
func (s *Server) handleListNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "nodes endpoint stub"}`))
}

func (s *Server) Run(ctx context.Context, addr string) error {
	httpServer := &http.Server{Addr: addr, Handler: s.mux}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		httpServer.Shutdown(shutdownCtx)
	}()

	log.Printf("listening on %s", addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
