package links

import (
	"net/http"
	"strconv"
	"test/packages/request"
	"test/packages/response"

	"gorm.io/gorm"
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
		for {

			existedLink, _ := handler.LinksRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinksRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinksHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinksUpdateRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		idInt, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.LinksRepository.Update(&Link{
			Model: gorm.Model{ID: uint(idInt)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, link, http.StatusOK)
	}
}

func (handler *LinksHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		idInt, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.LinksRepository.GetById(uint(idInt))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		handler.LinksRepository.Delete(uint(idInt))

		response.Json(w, nil, http.StatusOK)
	}
}

func (handler *LinksHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinksRepository.GetByHash(hash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
