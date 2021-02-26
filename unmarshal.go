package unmarshal

import (
	"github.com/buger/jsonparser"
)

type ConversionError struct {
	from string
	to   string
}

func (e ConversionError) Error() string {
	return "can not convert from " + e.from + " to " + e.to
}

func newError(from jsonparser.ValueType, to string) ConversionError {
	return ConversionError{
		from: from.String(),
		to:   to,
	}
}
