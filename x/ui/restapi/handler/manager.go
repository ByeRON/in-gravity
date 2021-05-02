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
	if err := m.Operator.SetupService(); err != nil {
		return
	}
	if err := m.Operator.ParseInput(); err != nil {
		return
	}
	m.Operator.ExecuteService()
}
