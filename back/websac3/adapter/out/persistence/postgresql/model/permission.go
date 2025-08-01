package model

type Permission struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ModuleID uint   `gorm:"not null" json:"module_id"`
	Module   Module `gorm:"foreignKey:ModuleID;references:ID"`
	ActionID uint   `gorm:"not null" json:"action_id"`
	Action   Action `gorm:"foreignKey:ActionID;references:ID"`
	Roles    []Role `gorm:"many2many:role_permissions"`
}
