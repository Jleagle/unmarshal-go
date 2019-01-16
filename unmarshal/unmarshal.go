package unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(in []byte, out interface{}) (err error) {

	// Unmarshal into map
	m := map[string]interface{}{}
	err = json.Unmarshal(in, &m)

	// Fix types
	m = fix(m, reflect.TypeOf(out))

	// Marshal back into bytes
	b, err := json.Marshal(m)

	// Unmarshal into struct
	return json.Unmarshal(b, out)
}

func fix(source map[string]interface{}, destinationType reflect.Type) (ret map[string]interface{}) {

	// get underlying type from ptr
	if destinationType.Kind() == reflect.Ptr {
		destinationType = destinationType.Elem()
	}

	// If source is empty
	if reflect.DeepEqual(source, reflect.Zero(reflect.TypeOf(source)).Interface()) {
		fmt.Println("empty source")
		return nil
	}

	if destinationType.Kind() == reflect.Struct && source == nil {
		fmt.Println("empty source")
		return nil
	}

	for i := 0; i < destinationType.NumField(); i++ {

		destinationField := destinationType.Field(i)

		fieldName := getJsonKey(destinationField)

		sourceVal := source[fieldName]
		if sourceVal == nil {
			continue
		}

		srcValKind := reflect.TypeOf(sourceVal).Kind()
		destinationFieldKind := destinationField.Type.Kind()

		switch destinationFieldKind {
		case reflect.String:

			break

		case reflect.Bool:

			source[fieldName] = toBool(srcValKind, destinationFieldKind, sourceVal, fieldName)
			break

		case reflect.Float64:

			source[fieldName] = toFloat(srcValKind, destinationFieldKind, sourceVal, fieldName)
			break

		default:
			errLog(srcValKind, destinationFieldKind, fieldName)
		}
	}

	return source
}

func toFloat(srcValKind reflect.Kind, destinationFieldKind reflect.Kind, srcVal interface{}, fieldName string) float64 {

	switch srcValKind {

	case reflect.String:

		f, err := strconv.ParseFloat(srcVal.(string), 64)
		if err != nil {
			fmt.Println(err)
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

		errLog(srcValKind, destinationFieldKind, fieldName)
		return 0
	}
}

func toBool(srcValKind reflect.Kind, destinationFieldKind reflect.Kind, srcVal interface{}, fieldName string) interface{} {

	switch srcValKind {
	case reflect.String:

		b, err := strconv.ParseBool(srcVal.(string))
		if err != nil {
			fmt.Println(err)
		}

		return b

	case reflect.Int | reflect.Int64 | reflect.Float64:

		return srcVal != 0

	case reflect.Bool:

		return srcVal

	default:

		errLog(srcValKind, destinationFieldKind, fieldName)
		return nil
	}
}

func getJsonKey(field reflect.StructField) (key string) {

	tag := field.Tag.Get("json")
	if tag == "" {

		key = field.Name

	} else {

		commaIndex := strings.Index(tag, ",")
		if commaIndex > 0 {
			key = tag[:commaIndex]
		} else {
			key = tag
		}
	}

	return key
}

func errLog(srcValKind reflect.Kind, destinationFieldKind reflect.Kind, fieldName string) {

	log.Printf("Unable to convert %s to %s (%s)", srcValKind, destinationFieldKind, fieldName)
}
