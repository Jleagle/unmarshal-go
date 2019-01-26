package main

import (
	"errors"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
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
type CFloat float64

func (i *CFloat) UnmarshalJSON(b []byte) error {

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
			*i = CFloat(j)

		} else {

			k, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*i = CFloat(k)

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

// type ScStringSlice []string
//
// func (s *ScStringSlice) UnmarshalJSON(b []byte) error {
// 	var raw, err = jsonparser.GetString(b)
// 	if err != nil {
// 		return err
// 	}
//
// 	if len(raw) == 0 {
// 		return nil
// 	}
//
// 	*s = ScStringSlice(strings.Split(raw, ","))
//
// 	return nil
// }
//
// type IntSlice []int
//
// func (is IntSlice) UnmarshalJSON(e *xml.Encoder, start xml.StartElement) error {
// 	intSliceString := ""
// 	if len(is) == 0 {
// 		return nil
// 	}
//
// 	b := make([]string, len(is))
// 	for i, v := range is {
// 		b[i] = strconv.Itoa(v)
// 	}
//
// 	intSliceString = strings.Join(b, ",")
// 	e.EncodeElement(intSliceString, start)
//
// 	return nil
// }
//
// type CBigInt big.Int
//
// func (is CBigInt) UnmarshalJSON(e *xml.Encoder, start xml.StartElement) error {
// 	intSliceString := ""
// 	if len(is) == 0 {
// 		return nil
// 	}
//
// 	b := make([]string, len(is))
// 	for i, v := range is {
// 		b[i] = strconv.Itoa(v)
// 	}
//
// 	intSliceString = strings.Join(b, ",")
// 	e.EncodeElement(intSliceString, start)
//
// 	return nil
// }
//
// type CBigFloat big.Float
//
// func (is CBigFloat) UnmarshalJSON(e *xml.Encoder, start xml.StartElement) error {
// 	intSliceString := ""
// 	if len(is) == 0 {
// 		return nil
// 	}
//
// 	b := make([]string, len(is))
// 	for i, v := range is {
// 		b[i] = strconv.Itoa(v)
// 	}
//
// 	intSliceString = strings.Join(b, ",")
// 	e.EncodeElement(intSliceString, start)
//
// 	return nil
// }
//
// type CTime time.Time
//
// func (i *CTime) UnmarshalJSON(b []byte) error {
//
// 	var value, dataType, _, err = jsonparser.Get(b)
// 	if err != nil {
// 		return err
// 	}
//
// 	switch dataType {
// 	case jsonparser.String, jsonparser.Number, jsonparser.Boolean, jsonparser.Array:
//
// 		*i = CString(value)
// 		return nil
//
// 	// case jsonparser.Array:
// 	//
// 	// 	var sli []string
// 	//
// 	// 	_, err = jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset int, err error) {
// 	// 		sli = append(sli, string(value2))
// 	// 	})
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	//
// 	// 	*i = CString(value)
// 	// 	return nil
//
// 	default:
//
// 		return errors.New("can not convert: " + types[dataType] + " to bool")
// 	}
// }
