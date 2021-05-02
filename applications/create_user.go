package applications

type CreateUserPresenterInterface interface {
	Output(CreateUserOutput)
	OutputError(error)
}

type CreateUserInput struct {
	Name string
}

type CreateUserOutput struct {
	UserId string
}

type CreateUserApplicationServiceInterface interface {
	//Handle(CreateUserInput)
	Handle()
}

type CreateUserApplicationService struct {
	presenter CreateUserPresenterInterface
}

func NewCreateUserApplicationService(
	presenter CreateUserPresenterInterface,
) CreateUserApplicationService {
	return CreateUserApplicationService{
		presenter: presenter,
	}
}

func (s CreateUserApplicationService) Handle() {
	s.presenter.Output(
		CreateUserOutput{
			UserId: "hogehoge",
		},
	)
}
