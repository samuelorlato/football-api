package ports

type Validator interface {
	Struct(obj interface{}) error
	GetErrors(err error, input interface{}) interface{}
}
