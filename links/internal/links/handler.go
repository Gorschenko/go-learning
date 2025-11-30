package links

import (
	"fmt"
	"net/http"
)

type LinksHandlerDeps struct {
	LinksRepository *LinksRepository
}

type LinksHandler struct {
	LinksRepository *LinksRepository
}

func NewLinksHandler(router *http.ServeMux, deps LinksHandlerDeps) {
	handler := &LinksHandler{
		LinksRepository: deps.LinksRepository,
	}
	router.HandleFunc("POST /links", handler.Create())
	router.HandleFunc("PATCH /links/{id}", handler.Update())
	router.HandleFunc("DELETE /links/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
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
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func (handler *LinksHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
