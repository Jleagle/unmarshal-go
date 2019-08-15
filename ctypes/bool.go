package ctypes

import (
	"errors"
	"strconv"

	"github.com/buger/jsonparser"
)

type CBool bool

func (i *CBool) UnmarshalJSON(b []byte) error {

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
	case jsonparser.String, jsonparser.Number, jsonparser.Boolean:

		b, _ := strconv.ParseBool(str)
		*i = CBool(b)

		return nil

	case jsonparser.Null:

		*i = false
		return nil

	}

	return errors.New("can not convert " + dataType.String() + " to bool")
}
