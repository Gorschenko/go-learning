package cars

import (
	"net/http"
)

func NewCarsHandler(router *http.ServeMux, dependencies *CarsHandlerDependencies) {
	// handler := &CarsHandler{
	// 	CarsService: dependencies.CarsService,
	// }

	// addCarURL := cars_api.AddCarMethod + " " + cars_api.AddCarPath
	// router.Handle(
	// 	addCarURL,
	// 	middlewares.ValidateBody[cars_api.AddCarRequestBodyDto](handler.AddCar()),
	// )
}

// func (h *CarsHandler) AddCar() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {}
// }
