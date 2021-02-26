package unmarshal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
)

type String string

func (i *String) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = ""
		return nil
	}

	str := string(data)

	switch dataType {
	case jsonparser.Object:

		if str == "{}" {
			str = ""
		}

		*i = String(fmt.Sprint(str))
		return nil

	case jsonparser.String, jsonparser.Number, jsonparser.Boolean:

		*i = String(str)
		return nil

	case jsonparser.Null:

		*i = ""
		return nil

	case jsonparser.Array:

		var slice []string
		_, err = jsonparser.ArrayEach(data, func(value2 []byte, dataType2 jsonparser.ValueType, offset int, err error) {
			slice = append(slice, string(value2))
		})
		if err != nil {
			return err
		}

		*i = String(strings.Join(slice, ","))
		return nil
	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}
