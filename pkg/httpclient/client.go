package httpclient

import (
	"net/http"
	"time"
)

func NewBearerClient(token string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			underlying: http.DefaultTransport,
			token:      token,
		},
		Timeout: 15 * time.Second,
	}
}

type authTransport struct {
	underlying http.RoundTripper
	token      string
}

func (a *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.token)
	return a.underlying.RoundTrip(req)
}
