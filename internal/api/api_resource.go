package api

import "net/http"

type ApiRoute struct {
	Url     string
	Handler http.HandlerFunc
	Method  string
}

type ApiResource interface {
	GetRoutes() []ApiRoute
}
