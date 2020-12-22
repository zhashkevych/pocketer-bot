package server

import (
	"net/http"
)

type RedirectServer struct {
	server      *http.Server
	redirectUrl string
}

func NewRedirectServer(redirectUrl string) *RedirectServer {
	return &RedirectServer{redirectUrl: redirectUrl}
}

func (s *RedirectServer) Start() error {
	s.server = &http.Server{
		Handler: s,
		Addr:    ":80",
	}

	return s.server.ListenAndServe()
}

func (s *RedirectServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", s.redirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)
}
