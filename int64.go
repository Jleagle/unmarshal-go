package unmarshal

import (
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

type Int64 int64

func (i *Int64) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = 0
		return nil
	}

	str := string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number:

		if strings.Contains(str, ".") {

			j, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return err
			}
			*i = Int64(j)

		} else {

			k, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*i = Int64(k)
		}

		return nil

	case jsonparser.Boolean:

		b, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}

		if b {
			*i = 1
		} else {
			*i = 0
		}

		return nil

	case jsonparser.Null:

		*i = 0

		return nil
	}

	return newError(dataType, "int64")
}
