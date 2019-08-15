package ctypes

import (
	"errors"
	"time"

	"github.com/buger/jsonparser"
)

type CTime time.Time

func (i *CTime) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CTime{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.Number, jsonparser.String:

		var t time.Time
		strLen := len(str)

		var dur time.Duration
		if strLen < 12 {
			dur, err = time.ParseDuration(str + "s") // Second
		} else if strLen < 15 {
			dur, err = time.ParseDuration(str + "ms") // Milli
		} else if strLen < 18 {
			dur, err = time.ParseDuration(str + "us") // Micro
		} else {
			dur, err = time.ParseDuration(str + "ns") // Nano
		}

		if err != nil {
			return err
		}

		t.Add(dur)

		*i = CTime(t)
		return nil

	case jsonparser.Null:

		*i = CTime{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}
