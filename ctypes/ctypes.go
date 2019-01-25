package ctypes

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/bugsnag/bugsnag-go/errors"
	"strconv"
)

type CString string

func (i *CString) UnmarshalJSON(b []byte) error {

	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	switch typex {
	case jsonparser.String:

		return srcVal.(string), err

	case jsonparser.Number:

		return strconv.Itoa(srcVal.(int)), err
		return strconv.FormatInt(srcVal.(int64), 10), err
		return strconv.FormatFloat(srcVal.(float64), 'f', -1, 64), err

	case jsonparser.Object:
		fmt.Println("OS X.")
	case jsonparser.Array:
		fmt.Println("OS X.")
	case jsonparser.Boolean:

		return strconv.FormatBool(srcVal.(bool)), err

	case jsonparser.Null:
		fmt.Println("Linux.")
	default:
		return errors.New("can not convert " + strconv.Itoa(typex) + " to bool")
	}

	return nil
}

type CInt int

func (i *CInt) UnmarshalJSON(b []byte) error {

	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	switch typex {
	case jsonparser.String:

		s := srcVal.(string)
		if s == "" {
			return 0, err
		}

		return strconv.Atoi(s)

	case jsonparser.Number:

		return srcVal.(int), err

		return int(srcVal.(float64)), err

	case jsonparser.Object:
		fmt.Println("OS X.")
	case jsonparser.Array:
		fmt.Println("OS X.")
	case jsonparser.Boolean:

		if srcVal.(bool) {
			return 1, err
		}
		return 0, err

	case jsonparser.Null:
		fmt.Println("Linux.")
	default:
		return errors.New("can not convert " + strconv.Itoa(typex) + " to bool")
	}

	return nil
}

type CBool bool

func (t *CBool) UnmarshalJSON(b []byte) error {

	var data, typex, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	switch typex {
	case jsonparser.String:

		b, _ := strconv.ParseBool(data.(string))
		return b, nil

	case jsonparser.Number:

		return data != 0, err

	case jsonparser.Object:
		fmt.Println("OS X.")
	case jsonparser.Array:
		fmt.Println("OS X.")
	case jsonparser.Boolean:

		return data.(bool), err

	case jsonparser.Null:
		fmt.Println("Linux.")
	default:
		return errors.New("can not convert " + strconv.Itoa(typex) + " to bool")
	}

	fmt.Println(string(b))
	fmt.Println(typex)

	*t = CBool("1" == string(raw))

	return nil
}

type CFloat float64

func (f *CFloat) UnmarshalJSON(b []byte) error {

	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	switch typex {
	case jsonparser.String:

		s := srcVal.(string)
		if s == "" {
			return 0, err
		}

		return strconv.ParseFloat(s, 64)

	case jsonparser.Number:

		return float64(srcVal.(int)), err

		return float64(srcVal.(int64)), err

		return srcVal.(float64), err

		if srcVal.(bool) {
			return 1, err
		}
		return 0, err

	case jsonparser.Object:
		fmt.Println("OS X.")
	case jsonparser.Array:
		fmt.Println("OS X.")
	case jsonparser.Boolean:
		fmt.Println("Linux.")
	case jsonparser.Null:
		fmt.Println("Linux.")
	default:
		return errors.New("can not convert " + strconv.Itoa(typex) + " to bool")
	}

	return nil
}
