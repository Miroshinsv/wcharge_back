package usecase

// UseCase controller -> UseCase -> repo -> entity
type UseCase struct {
	postgres PostgresRepo
	//webapi  WebAPIRepo
	//mqtt    MQTTApiRepo
}

func New(r PostgresRepo) *UseCase {
	return &UseCase{
		postgres: r,
	}
}
