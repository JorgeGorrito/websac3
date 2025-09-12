package entity

type EmailAttachment struct {
	ID          uint
	Filename    string
	ContentType string
	Data        []byte
	Size        int64
}


