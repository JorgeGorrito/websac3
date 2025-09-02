package entity

type CourseNature struct {
	ID    uint
	Names []CourseNatureName
}

type CourseNatureName struct {
	ID             uint
	Lang           string
	Name           string
	CourseNatureID uint
}
