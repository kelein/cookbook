package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

// Application Key and Secret
const (
	AppKey    = "govoyage"
	AppSecret = "go-grpc-voyage"
)

// AccessToken for request authentication
type AccessToken struct {
	Token string `json:"token"`
}

// API access request invoke
type API struct{ URL string }

// NewAPI creates a new API instance
func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("auth?app-key=%sapp-secret=%s", AppKey, AppSecret)
	body, err := a.get(ctx, url)
	if err != nil {
		return "", err
	}

	token := AccessToken{}
	if err := json.Unmarshal(body, &token); err != nil {
		slog.Error("failed to unmarshal access token", "error", err)
		return "", err
	}
	return token.Token, nil
}

func (a *API) get(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetTags return list of tags
func (a *API) GetTags(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		slog.Error("failed to get access token", "error", err)
		return nil, err
	}
	body, err := a.get(ctx, fmt.Sprintf(
		"/api/v1/tags?name=%s&token=%s",
		name, token,
	))
	if err != nil {
		return nil, err
	}
	return body, nil
}
