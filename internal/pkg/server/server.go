package server

import (
	"github.com/audetv/telegram-bot-getpocket/internal/db/bbolt/tokenstore"
	"github.com/audetv/telegram-bot-getpocket/internal/pkg/repos/token"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server          *http.Server
	pocketClient    *pocket.Client
	tokenRepository tokenstore.TokenRepository
	redirectURL     string
}

func NewAuthorizationServer(
	pocketClient *pocket.Client,
	tokenRepository tokenstore.TokenRepository,
	redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{
		pocketClient:    pocketClient,
		tokenRepository: tokenRepository,
		redirectURL:     redirectURL}
}

func (as *AuthorizationServer) Start() error {
	as.server = &http.Server{
		Addr:    ":80",
		Handler: as,
	}

	return as.server.ListenAndServe()
}

func (as *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := as.tokenRepository.Get(chatID, token.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authResp, err := as.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = as.tokenRepository.Create(chatID, authResp.AccessToken, token.AccessTokens)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", as.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
