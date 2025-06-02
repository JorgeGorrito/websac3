package command

type CreateAccessRequestCommand struct {
	Person CreatePersonCommand `mapper:"person"`
}
