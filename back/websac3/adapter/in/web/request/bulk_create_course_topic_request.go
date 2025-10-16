package request

import "mime/multipart"

type BulkCreateCourseTopicRequest struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
	CourseID uint                  `form:"course_id" binding:"required"`
}



