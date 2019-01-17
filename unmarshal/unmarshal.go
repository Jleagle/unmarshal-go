package unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Mutator func(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (interface{}, error)

var mutators = map[reflect.Kind]Mutator{}

func init() {
	mutators[reflect.Bool] = boolMutator
	mutators[reflect.Float64] = floatMutator
	mutators[reflect.Int] = intMutator
	mutators[reflect.Map] = mapMutator
	mutators[reflect.Ptr] = pointerMutator
	mutators[reflect.Slice] = sliceMutator
	mutators[reflect.String] = stringMutator
	mutators[reflect.Struct] = structMutator
}

//noinspection GoUnusedExportedFunction
func SetMutator(kind reflect.Kind, m Mutator) {
	mutators[kind] = m
}

//noinspection GoUnusedExportedFunction
func GetMutator(kind reflect.Kind) Mutator {
	return mutators[kind]
}

func Unmarshal(in []byte, out interface{}) (err error) {

	// Unmarshal into map
	m := map[string]interface{}{}
	err = json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	// Fix types
	m, err = mutate(m, reflect.TypeOf(out))
	if err != nil {
		return err
	}

	// Marshal back into bytes
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Unmarshal into struct
	return json.Unmarshal(b, out)
}

func mutate(source interface{}, destinationType reflect.Type) (ret map[string]interface{}, err error) {

	// get underlying type from ptr
	if destinationType.Kind() == reflect.Ptr {
		destinationType = destinationType.Elem()
	}

	// If source is empty
	if reflect.DeepEqual(source, reflect.Zero(reflect.TypeOf(source)).Interface()) {
		return nil, nil
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
				sourceMap[fieldName], err = mutator(destinationField.Type, sourceVal, fieldName)
			} else {
				err = errLog(sourceVal, destinationField.Type.Kind(), fieldName)
			}

			if err != nil {
				return sourceMap, err
			}
		}
	}

	return sourceMap, err
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

func errLog(sourceVal interface{}, destinationFieldKind reflect.Kind, fieldName string) error {

	srcValKind := reflect.TypeOf(sourceVal).Kind()

	return fmt.Errorf("unable to convert %s to %s (%s)", srcValKind, destinationFieldKind, fieldName)
}
