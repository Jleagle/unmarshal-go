package unmarshal

import (
	"errors"
	"math/big"

	"github.com/buger/jsonparser"
)

type BigFloat big.Float

func (i *BigFloat) UnmarshalJSON(b []byte) error {

	var data, dataType, _, err = jsonparser.Get(b)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		*i = BigFloat{}
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

		*i = BigFloat(*bigFloat)
		return nil

	case jsonparser.Null:

		*i = BigFloat{}
		return nil
	}

	return newError(dataType, "big.Float")
}
