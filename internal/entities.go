package internal

import (
	"net/http"
	"time"
)

type Client struct {
	url string
	port int
	timeout time.Duration
	httpClient *http.Client
}
