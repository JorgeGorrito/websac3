package entity

type CourseType struct {
	ID    uint
	Names []CourseTypeName
}

type CourseTypeName struct {
	ID          uint
	Lang        string
	Name        string
	CourseTypeID uint
}
