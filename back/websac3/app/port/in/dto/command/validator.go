package command

type Validator interface {
	Validate() error
}
