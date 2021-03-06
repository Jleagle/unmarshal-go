package unmarshal

import (
	"strconv"

	"github.com/buger/jsonparser"
)

type Bool bool

func (i *Bool) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = false
		return nil
	}

	str := string(data)

	switch dataType {
	case jsonparser.Object:

		if str == "{}" {
			*i = false
		} else {
			*i = true
		}
		return nil

	case jsonparser.String, jsonparser.Number, jsonparser.Boolean:

		b, _ := strconv.ParseBool(str)
		*i = Bool(b)

		return nil

	case jsonparser.Null:

		*i = false
		return nil
	}

	return newError(dataType, "bool")
}
