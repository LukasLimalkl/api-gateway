package entities

import (
	"net/http"
	"time"
)

type Client struct {
	url string
	timeout time.Duration
	htppClient *http.Client
}