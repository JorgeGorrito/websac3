package entity

type Permission struct {
	ID       uint
	ModuleID uint
	Module   *Module
	ActionID uint
	Action   *Action
}

func (p *Permission) String() string {
	var moduleName string
	var actionName string

	if p.Module != nil {
		moduleName = p.Module.Name
	}
	if p.Action != nil {
		actionName = p.Action.Name
	}

	return moduleName + ":" + actionName
}
