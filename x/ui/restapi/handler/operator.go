package handler

type RequestHandleOperatorInterface interface {
	SetupPresenter() error
	//ParseInput() error
	//SetupService() error
	ExecuteService()
}
