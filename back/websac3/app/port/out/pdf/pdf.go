package pdf

type Converter interface {
	HTMLToPDF(html string) ([]byte, error)
}
