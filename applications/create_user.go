package applications

import (
	"in-gravity/domains/entities"
	"in-gravity/domains/vo"
)

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
	Handle(CreateUserInput)
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

func (s CreateUserApplicationService) Handle(input CreateUserInput) {

	userId, err := vo.NewUserId()
	if err != nil {
		s.presenter.OutputError(err)
	}

	userName, err := vo.NewUserName(input.Name)
	if err != nil {
		s.presenter.OutputError(err)
	}

	entities.NewUser(userId, userName)

	s.presenter.Output(
		CreateUserOutput{
			UserId: userId.String(),
		},
	)
}
