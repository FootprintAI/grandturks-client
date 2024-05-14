package format

import "io"

type Header interface {
	Header() []string
}

type Rower interface {
	Rows() [][]string
}

type HeaderAndRows interface {
	Header
	Rower
}

var DefaultTypeFormatter TypeFormat = TypeFormatTable

func NewDefaultFormatter(hnr HeaderAndRows) *DefaultFormatter {
	return NewDefaultFormatterWithFormat(hnr, DefaultTypeFormatter)
}

func NewDefaultFormatterWithFormat(hnr HeaderAndRows, f TypeFormat) *DefaultFormatter {
	return &DefaultFormatter{
		f:   f,
		hnr: hnr,
	}
}

type DefaultFormatter struct {
	hnr HeaderAndRows
	f   TypeFormat
}

func (d *DefaultFormatter) Write(w io.Writer) error {
	fmtter := NewFormatWriter(newStructFormatter(d.f))
	fmtter.SetHeader(d.hnr.Header())
	fmtter.SetItems(d.hnr.Rows())
	return fmtter.Write(w)
}
