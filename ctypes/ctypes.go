package ctypes

import (
	"errors"
	"github.com/buger/jsonparser"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var types = map[jsonparser.ValueType]string{
	jsonparser.NotExist: "not-exist",
	jsonparser.String:   "string",
	jsonparser.Number:   "number",
	jsonparser.Object:   "object",
	jsonparser.Array:    "array",
	jsonparser.Boolean:  "bool",
	jsonparser.Null:     "null",
	jsonparser.Unknown:  "unknown",
}

type CString string

func (i *CString) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = ""
		return nil
	}

	switch dataType {
	case jsonparser.String, jsonparser.Number, jsonparser.Boolean, jsonparser.Array:

		*i = CString(data)
		return nil

	case jsonparser.Null:

		*i = ""

		// case jsonparser.Array:
		//
		// 	var sli []string
		//
		// 	_, err = jsonparser.ArrayEach(data, func(value2 []byte, dataType2 jsonparser.ValueType, offset int, err error) {
		// 		sli = append(sli, string(value2))
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}
		//
		// 	*i = CString(data)
		// 	return nil
	}

	return errors.New("can not convert: " + types[dataType] + " to bool")
}

type CInt int

func (i *CInt) UnmarshalJSON(b []byte) error {

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
			*i = CInt(j)

		} else {

			k, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*i = CInt(k)

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

	return errors.New("can not convert " + types[dataType] + " to int")
}

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

	return errors.New("can not convert " + types[dataType] + " to int64")
}

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

	}

	return errors.New("can not convert " + types[dataType] + " to bool")
}

//
type CFloat64 float64

func (i *CFloat64) UnmarshalJSON(b []byte) error {

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
			*i = CFloat64(j)

		} else {

			k, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*i = CFloat64(k)

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

	return errors.New("can not convert " + types[dataType] + " to float64")
}

type SStringSlice []string

func (i *SStringSlice) UnmarshalJSON(b []byte) error {

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

		*i = SStringSlice(strings.Split(str, ","))
		return nil

	case jsonparser.Number, jsonparser.Boolean:

		*i = []string{str}
		return nil

	case jsonparser.Null:

		*i = []string{}
		return nil

	}

	return errors.New("can not convert " + types[dataType] + " to string slice")
}

type SIntSlice []int

func (i *SIntSlice) UnmarshalJSON(b []byte) error {

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

	return errors.New("can not convert " + types[dataType] + " to string slice")
}

type CBigInt big.Int

type CBigFloat big.Float

type CTime time.Time
