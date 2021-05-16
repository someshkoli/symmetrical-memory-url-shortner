package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type response struct {
	Status bool   `json:"Status"`
	Data   string `json:"Data"`
}

// NewServer creates a new http server instance with given mux
func newServer(mux http.Handler) *http.Server {
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	return server
}

func loggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := log.Default()
			log.Printf("%s: %s", time.Now().String(), r.URL.String())
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
func respond(res http.ResponseWriter, data string, status bool, code int) {
	res.WriteHeader(code)
	rsp := response{
		Data:   data,
		Status: status,
	}
	json.NewEncoder(res).Encode(rsp)
}
