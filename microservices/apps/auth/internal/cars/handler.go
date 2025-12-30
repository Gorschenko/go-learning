package cars

import "net/http"

func NewCarsHandler(router *http.ServeMux, dependencies CarsHandlerDependencies) {
	_ = &CarsHandler{
		CarsService: dependencies.CarsService,
	}
}
