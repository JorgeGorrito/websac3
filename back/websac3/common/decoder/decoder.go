package decoder

type Decoder interface {
	Decode(filepath string, decoded any) error
}
