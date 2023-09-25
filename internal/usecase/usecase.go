package usecase

type UseCase struct {
	postgres PostgresRepo
	//webapi  WebAPI
	//mqtt    MQTTApi
}

func New(r PostgresRepo) *UseCase {
	return &UseCase{
		postgres: r,
	}
}
