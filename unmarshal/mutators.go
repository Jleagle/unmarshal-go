package unmarshal

import (
	"fmt"
	"reflect"
	"strconv"
)

func stringMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {

	case reflect.String:

		return srcVal.(string)

	case reflect.Int:

		return strconv.Itoa(srcVal.(int))

	case reflect.Int64:

		return strconv.FormatInt(srcVal.(int64), 10)

	case reflect.Float64:

		return strconv.FormatFloat(srcVal.(float64), 'f', 10, 64) // todo, make the precision an option

	case reflect.Bool:

		return strconv.FormatBool(srcVal.(bool))

	default:

		ErrLog(srcValKind, destinationFieldType.Kind(), fieldName)
		return 0
	}
}

func floatMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {

	case reflect.String:

		s := srcVal.(string)
		if s == "" {
			return 0
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		return f

	case reflect.Int:

		return float64(srcVal.(int))

	case reflect.Int64:

		return float64(srcVal.(int64))

	case reflect.Float64:

		return srcVal.(float64)

	case reflect.Bool:

		if srcVal.(bool) {
			return 1
		}
		return 0

	default:

		ErrLog(srcValKind, destinationFieldType.Kind(), fieldName)
		return 0
	}
}

func intMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.Bool:

		if srcVal.(bool) {
			return 1
		}
		return 0

	case reflect.Int:

		return srcVal.(int)

	case reflect.Float64:

		return int(srcVal.(float64))

	case reflect.String:

		s := srcVal.(string)
		if s == "" {
			return 0
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		return i

	default:

		ErrLog(srcValKind, destinationFieldType.Kind(), fieldName)
		return 0
	}
}

func boolMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.String:

		b, err := strconv.ParseBool(srcVal.(string))
		if err != nil {
			return false
		}

		return b

	case reflect.Int | reflect.Int64 | reflect.Float64:

		return srcVal != 0

	case reflect.Bool:

		return srcVal.(bool)

	default:

		ErrLog(srcValKind, destinationFieldType.Kind(), fieldName)
		return false
	}
}

func pointerMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	return mutate(srcVal, destinationFieldType.Elem())
}

func sliceMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{} {

	return nil
}
