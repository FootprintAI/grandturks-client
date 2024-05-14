package format

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TypeFormat string

func (t TypeFormat) String() string {
	return string(t)
}

const (
	TypeFormatUnknown TypeFormat = "unknown"
	TypeFormatTable   TypeFormat = "table"
	TypeFormatCSV     TypeFormat = "csv"
	TypeFormatJSON    TypeFormat = "json"
)

func MustParseFormat(s string) TypeFormat {
	switch s {
	case TypeFormatTable.String():
		return TypeFormatTable
	case TypeFormatCSV.String():
		return TypeFormatCSV
	case TypeFormatJSON.String():
		return TypeFormatJSON
	default:
		return TypeFormatUnknown
	}
}

func newStructFormatter(t TypeFormat) structFormatter {
	switch t {
	case TypeFormatTable:
		return structFormatFunc(tableStructFormatFunc)
	case TypeFormatCSV:
		return structFormatFunc(csvStructFormatFunc)
	case TypeFormatJSON:
		return structFormatFunc(jsonStructFormatFunc)
	}
	return structFormatFunc(noOpsFunc)
}

type structFormatter interface {
	Write(w io.Writer, headers []string, items [][]string) error
}

type structFormatFunc func(w io.Writer, headers []string, items [][]string) error

func (s structFormatFunc) Write(w io.Writer, headers []string, items [][]string) error {
	return s(w, headers, items)
}

func noOpsFunc(w io.Writer, headers []string, items [][]string) error {
	return errors.New("noops")
}

var (
	DefaultTableLintLength = 20
)

func tableStructFormatFunc(w io.Writer, headers []string, items [][]string) error {
	table := tablewriter.NewWriter(w)
	table.SetHeader(headers)
	for _, item := range items {
		table.Append(lint(item))
	}
	table.Render()
	return nil
}

func lint(s []string) []string {
	var linted []string
	for _, str := range s {
		linted = append(linted, limit(str, DefaultTableLintLength))
	}
	return linted
}

func limit(b string, length int) string {
	if len(b) > length {
		return fmt.Sprintf("%s...", b[:length])
	}
	return b
}

func csvStructFormatFunc(w io.Writer, headers []string, items [][]string) error {
	ww := csv.NewWriter(w)
	ww.Write(headers)
	for _, item := range items {
		ww.Write(item)
	}
	ww.Flush()
	return ww.Error()
}

func jsonStructFormatFunc(w io.Writer, headers []string, items [][]string) error {
	mapSlice := []map[string]string{}

	for _, item := range items {
		m := map[string]string{}
		for fieldIndex, fieldValue := range item {
			headerField := headers[fieldIndex]
			m[headerField] = fieldValue
		}
		mapSlice = append(mapSlice, m)
	}
	enc := json.NewEncoder(w)
	return enc.Encode(mapSlice)
}

func NewFormatWriter(formatter structFormatter) *FormatWriter {
	return &FormatWriter{
		formatter: formatter,
		headers:   []string{},
		items:     [][]string{},
	}
}

type FormatWriter struct {
	formatter structFormatter
	headers   []string
	items     [][]string
}

func (f *FormatWriter) SetHeader(headers []string) {
	f.headers = headers
}

func (f *FormatWriter) SetItems(items [][]string) {
	f.items = items
}

func (f *FormatWriter) Write(w io.Writer) error {
	return f.formatter.Write(w, f.headers, f.items)
}

func formatTime(t time.Time) string {
	return t.String()
}

func formatSizeInBytes(s int64) string {
	return byteCountBinary(s)
}

func byteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
