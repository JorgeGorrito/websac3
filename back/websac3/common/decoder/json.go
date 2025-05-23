package decoder

import (
	"encoding/json"
	"os"
)

type jsonDecoder struct{}

func Json() Decoder {
	return &jsonDecoder{}
}

func (j *jsonDecoder) Decode(filepath string, decoded any) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(decoded); err != nil {
		return nil
	}

	return nil
}
