package cars_api

import "pkg/database"

const (
	AddCarPath   = "/cars"
	AddCarMethod = "POST"
)

type AddCarRequestBodyDto struct {
	UserID    int    `json:"userID" validate:"required"`
	UserName  string `json:"userName"`
	ModelName string `json:"modelName"`
	Brand     string `json:"brand"`
	Color     string `json:"color"`
}

type AddCarResponseBodyDto struct {
	*database.Car `json:"car"`
}
