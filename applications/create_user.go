package applications

import (
	"in-gravity/domains/entities"
	"in-gravity/domains/repositories"
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
	presenter      CreateUserPresenterInterface
	userRepository repositories.UserRepositoryInterface
}

func NewCreateUserApplicationService(
	presenter CreateUserPresenterInterface,
	userRepository repositories.UserRepositoryInterface,
) CreateUserApplicationServiceInterface {
	return CreateUserApplicationService{
		presenter:      presenter,
		userRepository: userRepository,
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

	User := entities.NewUser(userId, userName)
	err = s.userRepository.Save(User)

	s.presenter.Output(
		CreateUserOutput{
			UserId: userId.String(),
		},
	)
}
