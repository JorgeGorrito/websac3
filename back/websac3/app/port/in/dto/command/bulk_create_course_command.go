package command

import "mime/multipart"

type BulkCreateCourseCommand struct {
	File            *multipart.FileHeader
	DegreeProgramID uint
	CreatedBy       uint
	Permissions     []string
}





