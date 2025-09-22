package request

type DeleteCourseRequest struct {
	CourseID uint `json:"course_id" binding:"required" validate:"required,min=1"`
}

