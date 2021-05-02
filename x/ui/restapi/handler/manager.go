package handler

type RequestHandleManagerInterface interface {
	Handle()
}

type RequestHandleManager struct {
	Operator RequestHandleOperatorInterface
}

func (m RequestHandleManager) Handle() {
	//SetupPresenter
	//ParseInput
	//SetupService
	m.Operator.ExecuteService()
}
