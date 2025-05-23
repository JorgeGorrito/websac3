package command

type CreateAccessRequestCommand struct {
	Person CreatePersonCommand `mapper:"person"`
}

func (c *CreateAccessRequestCommand) Validate() error {
	return c.Person.Validate()
}
