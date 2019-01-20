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

	// Unmarshal into generic JSON map
	m := map[string]interface{}{}
	err = json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	// Fix types
	m, err = mutate(m, out)
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

func mutate(source interface{}, destination interface{}) (map[string]interface{}, error) {

	var err error
	var destinationType = reflect.TypeOf(destination)
	var destinationKind = destinationType.Kind()
	var destinationValue = reflect.ValueOf(destination)

	// Get the real type of pointers
	if destinationKind == reflect.Ptr {
		destinationType = destinationType.Elem()
		destinationKind = destinationType.Kind()
		destinationValue = destinationValue.Elem()
	}

	// If source is empty
	if reflect.DeepEqual(source, reflect.Zero(reflect.TypeOf(source)).Interface()) {
		return nil, nil
	}

	sourceMap, sourceIsMap := source.(map[string]interface{})
	if sourceIsMap {

		switch k := destinationKind; k {
		case reflect.Struct:

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
			}

		case reflect.Map:

			var destinationMapValueType = reflect.TypeOf(destination).Elem() // map[string]main.DestinationData

			for _, key := range reflect.ValueOf(source).MapKeys() {

				value := reflect.ValueOf(source).MapIndex(key)

				out, err := mutate(value, destinationMapValueType)
				if err != nil {
					return sourceMap, err
				}

				fmt.Println(value)
				fmt.Println(out)

				destinationValue.SetMapIndex(key, reflect.ValueOf(out))
			}

		default:
			return sourceMap, fmt.Errorf("can not marshal into a " + k.String())
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
