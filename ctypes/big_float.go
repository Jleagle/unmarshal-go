package ctypes

import (
	"errors"
	"math/big"

	"github.com/buger/jsonparser"
)

type CBigFloat big.Float

func (i *CBigFloat) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CBigFloat{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number:

		var bigFloat *big.Float
		bigFloat, success := bigFloat.SetString(str)
		if success || bigFloat == nil {
			return errors.New("bigFloat.SetString error")
		}

		*i = CBigFloat(*bigFloat)
		return nil

	case jsonparser.Null:

		*i = CBigFloat{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}
