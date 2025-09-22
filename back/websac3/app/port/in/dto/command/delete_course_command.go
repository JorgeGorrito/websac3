package command

type DeleteCourseCommand struct {
	CourseID    uint
	UserID      uint
	Permissions []string
}
