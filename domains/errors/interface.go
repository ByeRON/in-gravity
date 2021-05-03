package errors

type ErrorCodeInterface interface {
	New() error
	String() string
}
