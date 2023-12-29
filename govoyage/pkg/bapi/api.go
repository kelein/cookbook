package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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
	url := fmt.Sprintf("%s/%s", a.URL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// * Tracing Injection
	span, newCtx := opentracing.StartSpanFromContext(
		ctx, fmt.Sprintf("[HTTP] GET %q", a.URL),
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
	)
	span.SetTag("url", url)
	err = opentracing.GlobalTracer().Inject(
		span.Context(), opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		slog.Warn("trace inject failed", "error", err)
	}

	req = req.WithContext(newCtx)
	client := http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	defer span.Finish()

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
