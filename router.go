package goapi

import (
	"fmt"
	"net/http"
)

type RouteConfig struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Router struct {
	routeTree Tree
}

func (r *Router) AddRoute(path string, method string, fn http.HandlerFunc) {
	r.routeTree.AddNode(path, &RouteConfig{
		Path:        path,
		Method:      method,
		HandlerFunc: fn,
	})
}

func (r *Router) FindByPath(path string, method string) (*RouteConfig, error) {
	node, err := r.routeTree.FindByPath(path)

	if err != nil {
		return nil, err
	}

	if node.configs[method] != nil {
		return node.configs[method], nil
	}

	return nil, fmt.Errorf("route not found")
}

func (r *Router) Get(path string, fn http.HandlerFunc) {
	r.AddRoute(path, "GET", wrapHandler(fn, "GET"))
}

func (r *Router) Post(path string, fn http.HandlerFunc) {
	r.AddRoute(path, "POST", wrapHandler(fn, "POST"))
}

func wrapHandler(fn http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			fn(w, r)
		}
	}
}
