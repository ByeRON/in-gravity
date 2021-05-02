package handler

type RequestHandleOperatorInterface interface {
	SetupPresenter() error
	SetupService() error
	ParseInput() error
	ExecuteService()
}
