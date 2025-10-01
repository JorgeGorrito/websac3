package command

import "mime/multipart"

type BulkCreateDegreeProgramCommand struct {
	File        *multipart.FileHeader
	CreatedBy   uint `validations:"required"`
	Permissions []string
}
