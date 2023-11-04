package httpApi

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

// just for testing CI, remove this comment asap
func GetHttpServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/1/ping", GetPingHandler())
	handler := handlers.LoggingHandler(os.Stdout, r)
	return &http.Server{
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
