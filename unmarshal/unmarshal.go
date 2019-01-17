package unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type mutator func(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) interface{}

var mutators = map[reflect.Kind]mutator{}

func init() {
	mutators[reflect.Bool] = boolMutator
	mutators[reflect.Float64] = floatMutator
	mutators[reflect.Int] = intMutator
	mutators[reflect.String] = stringMutator
	mutators[reflect.Ptr] = pointerMutator
	mutators[reflect.Slice] = sliceMutator
}

//noinspection GoUnusedExportedFunction
func SetMutator(kind reflect.Kind, m mutator) {
	mutators[kind] = m
}

func Unmarshal(in []byte, out interface{}) (err error) {

	// Unmarshal into map
	m := map[string]interface{}{}
	err = json.Unmarshal(in, &m)

	// Fix types
	m = mutate(m, reflect.TypeOf(out))

	// Marshal back into bytes
	b, err := json.Marshal(m)

	// Unmarshal into struct
	return json.Unmarshal(b, out)
}

func mutate(source interface{}, destinationType reflect.Type) (ret map[string]interface{}) {

	// get underlying type from ptr
	if destinationType.Kind() == reflect.Ptr {
		destinationType = destinationType.Elem()
	}

	// If source is empty
	if reflect.DeepEqual(source, reflect.Zero(reflect.TypeOf(source)).Interface()) {
		fmt.Println("empty source")
		return nil
	}

	sourceMap, isMap := source.(map[string]interface{})
	if isMap {
		for i := 0; i < destinationType.NumField(); i++ {

			destinationField := destinationType.Field(i)

			fieldName := getJsonKey(destinationField)
			sourceVal := sourceMap[fieldName]
			if sourceVal == nil {
				continue
			}

			mutator, exists := mutators[destinationField.Type.Kind()]
			if exists {
				sourceMap[fieldName] = mutator(destinationField.Type, sourceVal, fieldName)
			} else {
				ErrLog(sourceVal, destinationField.Type.Kind(), fieldName)
			}
		}
	}

	return sourceMap
}

func ErrLog(sourceVal interface{}, destinationFieldKind reflect.Kind, fieldName string) {

	srcValKind := reflect.TypeOf(sourceVal).Kind()
	log.Printf("Unable to convert %s to %s (%s)", srcValKind, destinationFieldKind, fieldName)
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
