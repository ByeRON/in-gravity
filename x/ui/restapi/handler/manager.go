package handler

type RequestHandleManagerInterface interface {
	Handle()
}

type RequestHandleManager struct {
	Operator RequestHandleOperatorInterface
}

func (m RequestHandleManager) Handle() {
	if err := m.Operator.SetupPresenter(); err != nil {
		return
	}
	//ParseInput
	if err := m.Operator.SetupService(); err != nil {
		return
	}
	m.Operator.ExecuteService()
}
