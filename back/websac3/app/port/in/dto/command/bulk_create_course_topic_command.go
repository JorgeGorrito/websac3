package command

import "mime/multipart"

type BulkCreateCourseTopicCommand struct {
	File        *multipart.FileHeader
	CourseID    uint
	UserID      uint
	Permissions []string
}






