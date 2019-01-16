package unmarshal

import (
	"encoding/json"
	"reflect"
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

	}

	if destinationType.Kind() == reflect.Struct && source == nil {
		return nil
	}

	for i := 0; i < destinationType.NumField(); i++ {

		destinationField := destinationType.Field(i)

		fieldName := getJsonKey(destinationField)

		srcVal := source[fieldName]
		if srcVal == nil {
			continue
		}

		//srcValKind := reflect.TypeOf(srcVal).Kind()

		switch destinationField.Type.Kind() {

		case reflect.String:

			break

		case reflect.Bool:

			break

		default:

		}
	}

	return source
}

func getJsonKey(field reflect.StructField) (key string) {

	tag := field.Tag.Get("json")
	if tag != "" {
		if commaIndex := strings.Index(tag, ","); commaIndex > 0 {
			key = tag[:commaIndex]
		} else {
			key = tag
		}
	} else {
		key = field.Name
	}

	return key
}
