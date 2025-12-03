package links

import (
	"fmt"
	"net/http"
	"strconv"
	"test/configs"
	"test/packages/middlewares"
	"test/packages/request"
	"test/packages/response"

	"gorm.io/gorm"
)

type LinksHandlerDeps struct {
	LinksRepository *LinksRepository
	Config          *configs.Config
}

type LinksHandler struct {
	LinksRepository *LinksRepository
}

func NewLinksHandler(router *http.ServeMux, deps LinksHandlerDeps) {
	handler := &LinksHandler{
		LinksRepository: deps.LinksRepository,
	}
	router.Handle("POST /links", middlewares.IsAuthenticated(handler.Create(), deps.Config))
	router.Handle("PATCH /links/{id}", middlewares.IsAuthenticated(handler.Update(), deps.Config))
	router.Handle("DELETE /links/{id}", middlewares.IsAuthenticated(handler.Delete(), deps.Config))
	router.Handle("GET /links", middlewares.IsAuthenticated(handler.GetAll(), deps.Config))
	router.Handle("GET /links/{hash}/goTo", middlewares.IsAuthenticated(handler.GoTo(), deps.Config))

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
		email, ok := r.Context().Value(middlewares.ContextEmailKey).(string)

		if ok {
			fmt.Println(email)
		}

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
		fmt.Println(hash)
		link, err := handler.LinksRepository.GetByHash(hash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinksHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitString := r.URL.Query().Get("limit")
		limitInt, err := strconv.Atoi(limitString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(limitInt)
	}
}
