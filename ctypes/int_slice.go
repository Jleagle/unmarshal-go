package ctypes

import (
	"errors"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

type IntSlice []int

func (i *IntSlice) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = []int{}
		return nil
	}

	str := string(data)

	switch dataType {
	case jsonparser.String:

		split := strings.Split(str, ",")
		var ints []int
		for _, v := range split {
			j, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			ints = append(ints, j)
		}

		*i = ints
		return nil

	case jsonparser.Number:

		j, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}

		*i = []int{int(j)}
		return nil

	case jsonparser.Boolean:

		j, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}

		if j {
			*i = []int{1}
		} else {
			*i = []int{0}
		}
		return nil

	case jsonparser.Null:

		*i = []int{}
		return nil

	}

	return errors.New("can not convert " + dataType.String() + " to string slice")
}
