package request

import "mime/multipart"

type BulkCreateCourseRequest struct {
	File            *multipart.FileHeader `form:"file" binding:"required"`
	DegreeProgramID uint                  `form:"degree_program_id" binding:"required"`
}








