package cars

func NewCarsService(dependencies *CarsServiceDependencies) *CarsService {
	return &CarsService{
		CarsRepository:  dependencies.CarsRepository,
		UsersRepository: dependencies.UsersRepository,
	}
}

// func (s *CarsService) AddCar(car *database.Car) *database.Car {
// 	existedUser, _ := s.UsersRepository.FindByEmail()
// }
