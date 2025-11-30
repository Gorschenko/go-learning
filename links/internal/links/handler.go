package links

import (
	"fmt"
	"net/http"
	"test/packages/request"
	"test/packages/response"
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
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinksCreateRequest](&w, r)

		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinksRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
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
