//+build wireinject

package registry

import (
	"in-gravity/applications"
	"in-gravity/infrastructures/mysql/repositories"

	"github.com/google/wire"
)

func ProvideCreateUserApplicationService(
	p applications.CreateUserPresenterInterface,
) (
	applications.CreateUserApplicationServiceInterface,
	error,
) {
	wire.Build(
		repositories.NewUserRepository,
		applications.NewCreateUserApplicationService,
	)
	return nil, nil
}
