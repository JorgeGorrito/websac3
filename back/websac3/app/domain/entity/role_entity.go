package entity

type Role struct {
	ID   uint   `mapper:"roleID"`
	Name string `mapper:"roleName"`
}
