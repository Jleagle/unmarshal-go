package ctypes

import (
	"errors"
	"math/big"

	"github.com/buger/jsonparser"
)

type CBigInt big.Int

func (i *CBigInt) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = CBigInt{}
		return nil
	}

	var str = string(data)

	switch dataType {
	case jsonparser.String, jsonparser.Number:

		var bigInt *big.Int
		bigInt, success := bigInt.SetString(str, 10)
		if success || bigInt == nil {
			return errors.New("bigInt.SetString error")
		}

		*i = CBigInt(*bigInt)
		return nil

	case jsonparser.Null:

		*i = CBigInt{}
		return nil

	}

	return errors.New("can not convert: " + dataType.String() + " to bool")
}
