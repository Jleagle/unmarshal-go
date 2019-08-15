package ctypes

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

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

	str := string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number, jsonparser.Boolean:

		*i = CString(str)
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

		*i = CString(strings.Join(slice, ","))
		return nil
	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
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

	return errors.New("can not convert " + dataType.String() + " to int")
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

	return errors.New("can not convert " + dataType.String() + " to int64")
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
		return nil

	}

	return errors.New("can not convert " + dataType.String() + " to bool")
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

	return errors.New("can not convert " + dataType.String() + " to float64")
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

	case jsonparser.Object:

		var slice []string
		err := jsonparser.ObjectEach(data, func(key, value2 []byte, valueType2 jsonparser.ValueType, offset int) error {
			slice = append(slice, string(value2))
			return nil
		})
		if err != nil {
			return err
		}

		*i = SStringSlice(slice)
		return nil
	}

	return errors.New("can not convert " + dataType.String() + " to string slice")
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

	return errors.New("can not convert " + dataType.String() + " to string slice")
}

type CBigInt big.Int

func (i *CBigInt) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CBigInt{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number:

		var bigInt *big.Int
		bigInt, success := bigInt.SetString(str, 10)
		if success || bigInt == nil {
			return errors.New("bigInt.SetString error")
		}

		*i = CBigInt(*bigInt)
		return nil

	case jsonparser.Null:

		*i = CBigInt{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}

type CBigFloat big.Float

func (i *CBigFloat) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CBigFloat{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number:

		var bigFloat *big.Float
		bigFloat, success := bigFloat.SetString(str)
		if success || bigFloat == nil {
			return errors.New("bigFloat.SetString error")
		}

		*i = CBigFloat(*bigFloat)
		return nil

	case jsonparser.Null:

		*i = CBigFloat{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}

type CTime time.Time

func (i *CTime) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CTime{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.Number, jsonparser.String:

		var t time.Time
		strLen := len(str)

		var dur time.Duration
		if strLen < 12 {
			dur, err = time.ParseDuration(str + "s") // Second
		} else if strLen < 15 {
			dur, err = time.ParseDuration(str + "ms") // Milli
		} else if strLen < 18 {
			dur, err = time.ParseDuration(str + "us") // Micro
		} else {
			dur, err = time.ParseDuration(str + "ns") // Nano
		}

		if err != nil {
			return err
		}

		t.Add(dur)

		*i = CTime(t)
		return nil

	case jsonparser.Null:

		*i = CTime{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}
