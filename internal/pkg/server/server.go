package server

import (
	"github.com/audetv/telegram-bot-getpocket/internal/db/bbolt/tokenstore"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/http"
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

}
