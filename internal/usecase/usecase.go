package usecase

type UseCase struct {
	postgres UserRepo
	//webapi  WebAPI
}

func New(r UserRepo) *UseCase {
	return &UseCase{
		postgres: r,
	}
}
