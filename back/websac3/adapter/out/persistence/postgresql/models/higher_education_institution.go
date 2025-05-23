package models

type HigherEducationInstitution struct {
	Snies uint `gorm:"primary_key"`

	SniesParent                *uint                       `gorm:"type:integer;null"`
	HigherEducationInstitution *HigherEducationInstitution `gorm:"foreignkey:SniesParent"`

	Name string `gorm:"type:varchar(128);not null"`

	OwnershipID uint      `gorm:"not null"`
	Ownership   Ownership `gorm:"foreignkey:OwnershipID"`

	InstitutionalCategoryID uint                  `gorm:"not null"`
	InstitutionalCategory   InstitutionalCategory `gorm:"foreignkey:InstitutionalCategoryID"`

	MunicipalityID uint         `gorm:"not null"`
	Municipality   Municipality `gorm:"foreignkey:MunicipalityID"`

	DepartmentID uint       `gorm:"not null"`
	Department   Department `gorm:"foreignkey:DepartmentID"`
}
