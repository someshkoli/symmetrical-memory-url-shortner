package api

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/someshkoli/symmetrical-memory-url-shorner/pkg/shortner"
	"github.com/someshkoli/symmetrical-memory-url-shorner/pkg/store"
)

type URLShortner struct {
	urlStore store.RecordStore
	server   *http.Server
	port     int
	host     string
}

func NewURLShortner(s store.RecordStore, port int, host string) *URLShortner {
	return &URLShortner{
		urlStore: s,
		port:     port,
		host:     host,
	}
}

func (h *URLShortner) registerRoutes() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc("/set", h.set)
	m.HandleFunc("/g/{*}", h.redirect)
	router := handlers.LoggingHandler(os.Stdout, m)
	return router
}

func (h *URLShortner) Start() error {
	h.server = newServer(h.registerRoutes())
	h.server.Addr = net.JoinHostPort("", strconv.Itoa(h.port))
	h.urlStore.Sync()
	return h.server.ListenAndServe()
}

func (h *URLShortner) set(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		respond(w, "url is required param", false, http.StatusBadRequest)
		return
	}

	short := shortner.NewPath(shortner.RandStringRunes(5))
	short, _ = h.urlStore.Set(url, short)

	shortenURL := shortner.NewURL(h.host, h.port, short)
	respond(w, shortenURL, true, http.StatusAccepted)
}

func (h *URLShortner) redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	orgURL, err := h.urlStore.Get(path)
	if err != nil {
		respond(w, "url not available in store", false, http.StatusAccepted)
		return
	}
	url, err := url.Parse(orgURL)
	if url.Scheme == "" {
		url.Scheme = "https"
	}
	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}
