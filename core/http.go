package core

import "net/http"

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)
