package models

type Status struct {
	ID   uint   `gorm:"primary_key" mapper:"statusID"`
	Name string `gorm:"type:varchar(32);not null" mapper:"statusName"`
}
