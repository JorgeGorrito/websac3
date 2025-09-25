package entity

type FormationLevel struct {
	ID   uint
	Name string
}

func (f *FormationLevel) IsRegistered() bool {
	return f.ID != 0
}
