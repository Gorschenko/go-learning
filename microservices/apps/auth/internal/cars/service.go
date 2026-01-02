package cars

func NewCarsService(dependencies *CarsServiceDependencies) *CarsService {
	return &CarsService{
		CarsRepository:  dependencies.CarsRepository,
		UsersRepository: dependencies.UsersRepository,
	}
}
