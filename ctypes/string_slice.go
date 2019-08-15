package ctypes

import (
	"errors"
	"strings"

	"github.com/buger/jsonparser"
)

type CStringSlice []string

func (i *CStringSlice) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = []string{}
		return nil
	}

	str := string(data)

	switch dataType {
	case jsonparser.String:

		*i = strings.Split(str, ",")
		return nil

	case jsonparser.Number, jsonparser.Boolean:

		*i = []string{str}
		return nil

	case jsonparser.Null:

		*i = []string{}
		return nil

	case jsonparser.Object:

		var slice []string
		err := jsonparser.ObjectEach(data, func(key, value2 []byte, valueType2 jsonparser.ValueType, offset int) error {
			slice = append(slice, string(value2))
			return nil
		})
		if err != nil {
			return err
		}

		*i = slice
		return nil
	}

	return errors.New("can not convert " + dataType.String() + " to string slice")
}
