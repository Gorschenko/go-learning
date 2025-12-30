package cars

type CarsHandlerDependencies struct {
	CarsService *CarsService
}

type CarsHandler struct {
	CarsService *CarsService
}

type CarsServiceDependencies struct {
	CarsRepository *CarsRepository
}

type CarsService struct {
	CarsRepository *CarsRepository
}
