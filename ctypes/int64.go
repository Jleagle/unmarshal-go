package ctypes

import (
	"errors"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

type CInt64 int64

func (i *CInt64) UnmarshalJSON(b []byte) error {

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
			*i = CInt64(j)

		} else {

			k, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*i = CInt64(k)

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

	return errors.New("can not convert " + dataType.String() + " to int64")
}
