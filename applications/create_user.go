package applications

type CreateUserPresenterInterface interface {
	Output()
	OutputError(error)
}
