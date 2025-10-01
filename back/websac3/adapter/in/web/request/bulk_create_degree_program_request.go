package request

import "mime/multipart"

type BulkCreateDegreeProgramRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
