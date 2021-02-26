package unmarshal

import (
	"errors"
	"math/big"

	"github.com/buger/jsonparser"
)

type BigInt big.Int

func (i *BigInt) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = BigInt{}
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

		*i = BigInt(*bigInt)
		return nil

	case jsonparser.Null:

		*i = BigInt{}
		return nil
	}

	return newError(dataType, "big.Int")
}
