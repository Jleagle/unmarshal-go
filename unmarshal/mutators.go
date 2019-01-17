package unmarshal

import (
	"gopkg.in/inf.v0"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func stringMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {

	case reflect.String:

		return srcVal.(string), err

	case reflect.Int:

		return strconv.Itoa(srcVal.(int)), err

	case reflect.Int64:

		return strconv.FormatInt(srcVal.(int64), 10), err

	case reflect.Float64:

		return strconv.FormatFloat(srcVal.(float64), 'f', -1, 64), err

	case reflect.Bool:

		return strconv.FormatBool(srcVal.(bool)), err
	}

	return "", errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

func floatMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {

	case reflect.String:

		s := srcVal.(string)
		if s == "" {
			return 0, err
		}

		return strconv.ParseFloat(s, 64)

	case reflect.Int:

		return float64(srcVal.(int)), err

	case reflect.Int64:

		return float64(srcVal.(int64)), err

	case reflect.Float64:

		return srcVal.(float64), err

	case reflect.Bool:

		if srcVal.(bool) {
			return 1, err
		}
		return 0, err
	}

	return 0, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

func intMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.Bool:

		if srcVal.(bool) {
			return 1, err
		}
		return 0, err

	case reflect.Int:

		return srcVal.(int), err

	case reflect.Float64:

		return int(srcVal.(float64)), err

	case reflect.String:

		s := srcVal.(string)
		if s == "" {
			return 0, err
		}

		return strconv.Atoi(s)
	}

	return 0, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

func boolMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.String:

		// If there is an error, just return false
		b, _ := strconv.ParseBool(srcVal.(string))
		return b, nil

	case reflect.Int | reflect.Int64 | reflect.Float64:

		return srcVal != 0, err

	case reflect.Bool:

		return srcVal.(bool), err
	}

	return false, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

//noinspection GoUnusedParameter
func pointerMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	return mutate(srcVal, destinationFieldType.Elem())
}

func sliceMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	slice := make([]interface{}, 0)

	switch srcValKind {
	case reflect.Map:

		for _, v := range srcVal.(map[string]interface{}) {
			slice = append(slice, v)
		}
		return slice, err

	case reflect.Slice:

		return srcVal, err

	case reflect.String:

		s := srcVal.(string)

		if strings.Contains(s, ",") {
			var ret []string
			for _, v := range strings.Split(s, ",") {
				ret = append(ret, strings.TrimSpace(v))
			}
			return ret, err
		}

		break
	}

	return slice, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

func mapMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.Map:

		return srcVal, err
	}

	m := map[interface{}]interface{}{}
	return m, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}

func structMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch destinationFieldType {
	case reflect.TypeOf(inf.Dec{}):

		d := &inf.Dec{}

		srcValKind := reflect.TypeOf(srcVal).Kind()

		switch srcValKind {

		case reflect.Float64 | reflect.String:

			if srcValKind == reflect.Float64 {
				srcVal = strconv.FormatFloat(srcVal.(float64), 'g', -1, 64)
			}

			var success bool
			d, success = d.SetString(srcVal.(string))
			if !success {
				return d, err
			}
		}

		return d, errLog(srcValKind, destinationFieldType.Kind(), fieldName)

	case reflect.TypeOf(time.Time{}):

		t := time.Time{}

		return t, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
	}

	return srcVal, errLog(srcValKind, destinationFieldType.Kind(), fieldName)
}
