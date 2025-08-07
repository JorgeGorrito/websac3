package entity

type HigherEducationInstitution struct {
	Snies uint

	HigherEducationInstitutionID     uint
	HigherEducationInstitutionParent *HigherEducationInstitution

	Name string

	OwnershipID uint
	Ownership   *Ownership

	InstitutionalCategoryID uint
	InstitutionalCategory   *InstitutionalCategory

	MunicipalityID uint
	Municipality   *Municipality

	DepartmentID uint
	Department   *Department
}
