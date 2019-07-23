package GODP

import (
	"net/http"
	"time"
)

type ODP struct {
	JWT         string
	BaseURL     string
	AppName     string
	ServiceName string
	Username    string
	Password    string
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

