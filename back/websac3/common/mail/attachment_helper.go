package mail

import (
	"io"
	"mime"
	"os"
	"path/filepath"
	"websac3/app/domain/entity"
)

// CreateAttachmentFromFile crea un adjunto a partir de un archivo en disco
func CreateAttachmentFromFile(filePath string) (*entity.EmailAttachment, error) {
	// Leer el archivo
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Obtener información del archivo
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	// Obtener nombre del archivo
	filename := filepath.Base(filePath)

	// Detectar tipo MIME
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &entity.EmailAttachment{
		Filename:    filename,
		ContentType: contentType,
		Data:        data,
		Size:        fileInfo.Size(),
	}, nil
}

// CreateAttachmentFromReader crea un adjunto a partir de un Reader
func CreateAttachmentFromReader(reader io.Reader, filename, contentType string) (*entity.EmailAttachment, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &entity.EmailAttachment{
		Filename:    filename,
		ContentType: contentType,
		Data:        data,
		Size:        int64(len(data)),
	}, nil
}

// CreateAttachmentFromBytes crea un adjunto a partir de datos en memoria
func CreateAttachmentFromBytes(data []byte, filename, contentType string) *entity.EmailAttachment {
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &entity.EmailAttachment{
		Filename:    filename,
		ContentType: contentType,
		Data:        data,
		Size:        int64(len(data)),
	}
}


