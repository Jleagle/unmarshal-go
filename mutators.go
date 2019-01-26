package main

import (
	"math/big"
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

	return "", nil
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

	return 0, nil
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

	return 0, nil
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

	return false, nil
}

//noinspection GoUnusedParameter
func pointerMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	return nil, nil
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

	return slice, nil
}

func mapMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch srcValKind {
	case reflect.Map:

		return srcVal, err
	}

	m := map[interface{}]interface{}{}
	return m, nil
}

func structMutator(destinationFieldType reflect.Type, srcVal interface{}, fieldName string) (i interface{}, err error) {

	srcValKind := reflect.TypeOf(srcVal).Kind()

	switch destinationFieldType {
	case reflect.TypeOf(big.Int{}):

		var bigInt = new(big.Int)

		if srcValKind == reflect.Int {
			bigInt = bigInt.SetInt64(int64(srcVal.(int)))
			return bigInt, err
		}

		if srcValKind == reflect.String {
			var success bool
			bigInt, success = bigInt.SetString(srcVal.(string), 10)
			if success {
				return bigInt, err
			}
		}

		return bigInt, nil

	case reflect.TypeOf(big.Float{}):

		var bigFloat = new(big.Float)

		if srcValKind == reflect.Float64 {
			bigFloat = bigFloat.SetFloat64(srcVal.(float64))
			return bigFloat, err
		}

		if srcValKind == reflect.String {
			var success bool
			bigFloat, success = bigFloat.SetString(srcVal.(string))
			if success {
				return bigFloat, err
			}
		}

		return bigFloat, nil

	case reflect.TypeOf(time.Time{}):

		t := time.Time{}

		if srcValKind == reflect.Float64 || srcValKind == reflect.Int || srcValKind == reflect.String {

			if srcValKind == reflect.Float64 {
				srcVal = strconv.FormatFloat(srcVal.(float64), 'g', -1, 64)
			}

			if srcValKind == reflect.Int {
				srcVal = strconv.FormatInt(int64(srcVal.(int)), 10)
			}

			timeStr := srcVal.(string)
			timeLen := len(timeStr)

			var dur time.Duration
			if timeLen < 12 {
				dur, err = time.ParseDuration(timeStr + "s") // Second
			} else if timeLen < 15 {
				dur, err = time.ParseDuration(timeStr + "ms") // Milli
			} else if timeLen < 18 {
				dur, err = time.ParseDuration(timeStr + "us") // Micro
			} else {
				dur, err = time.ParseDuration(timeStr + "ns") // Nano
			}

			t.Add(dur)

			return t, err
		}

		return t, nil
	}

	return srcVal, nil
}
