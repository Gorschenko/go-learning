package links

import (
	"net/http"
)

type LinksHandlerDeps struct {
}

type LinksHandler struct {
}

func NewLinksHandler(router *http.ServeMux, deps LinksHandlerDeps) {
	handler := &LinksHandler{}
	router.HandleFunc("POST /links", handler.Create())
	router.HandleFunc("PATCH /links/{id}", handler.Update())
	router.HandleFunc("DELETE /links/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

func (handler *LinksHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (handler *LinksHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinksHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinksHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
