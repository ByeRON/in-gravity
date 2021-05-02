package handler

type RequestHandleManagerInterface interface {
	Handle()
}

type RequestHandleManager struct {
	Operator RequestHandleOperatorInterface
}

func (m RequestHandleManager) Handle() {
<<<<<<< HEAD
	m.Operator.SetupPresenter()
=======
	if err := m.Operator.SetupPresenter(); err != nil {
		return
	}
>>>>>>> Add
	//ParseInput
	if err := m.Operator.SetupService(); err != nil {
		return
	}
	m.Operator.ExecuteService()
}
