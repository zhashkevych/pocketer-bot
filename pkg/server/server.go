package server

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhashkevych/go-pocket-sdk"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type AuthServer struct {
	server *http.Server
	logger *zap.Logger

	storage storage.TokenStorage
	client  *pocket.Client

	redirectUrl string
}

func NewAuthServer(redirectUrl string, storage storage.TokenStorage, client *pocket.Client) *AuthServer {
	return &AuthServer{
		redirectUrl: redirectUrl,
		storage:     storage,
		client:      client,
	}
}

func (s *AuthServer) Start() error {
	s.server = &http.Server{
		Handler: s,
		Addr:    ":80",
	}

	logger, _ := zap.NewDevelopment(zap.Fields(
		zap.String("app", "authorization server")))
	defer logger.Sync()

	s.logger = logger

	return s.server.ListenAndServe()
}

func (s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.logger.Debug("received unavailable HTTP method request",
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusForbidden)
		return
	}

	chatIDQuery := r.URL.Query().Get("chat_id")
	if chatIDQuery == "" {
		s.logger.Debug("received empty chat_id query param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDQuery, 10, 64)
	if err != nil {
		s.logger.Debug("received invalid chat_id query param",
			zap.String("chat_id", chatIDQuery))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.createAccessToken(r.Context(), chatID); err != nil {
		s.logger.Debug("failed to create access token",
			zap.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", s.redirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)
}

func (s *AuthServer) createAccessToken(ctx context.Context, chatID int64) error {
	requestToken, err := s.storage.Get(chatID, storage.RequestTokens)
	if err != nil {
		return errors.WithMessage(err, "failed to get request token")
	}

	authResp, err := s.client.Authorize(ctx, requestToken)
	if err != nil {
		return errors.WithMessage(err, "failed to authorize at Pocket")
	}

	if err := s.storage.Save(chatID, authResp.AccessToken, storage.AccessTokens); err != nil {
		return errors.WithMessage(err, "failed to save access token to storage")
	}

	return nil
}
