package handler

type RequestHandleManagerInterface interface {
	Handle()
}

type RequestHandleManager struct {
	Operator RequestHandleOperatorInterface
}

func (m RequestHandleManager) Handle() {
	m.Operator.SetupPresenter()
	//ParseInput
	//SetupService
	m.Operator.ExecuteService()
}
