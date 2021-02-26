package unmarshal

import (
	"encoding/json"
	"testing"
)

type SourceCData struct {
	StringFromInt     int
	StringFromFloat   float64
	StringFromBool    bool
	StringFromObject  map[string]interface{}
	StringFromObject2 map[string]interface{}

	BoolFromInt     int
	BoolFromFloat   float64
	BoolFromString  string
	BoolFromString2 string
	BoolFromString3 string
	BoolFromObject  map[string]interface{}
	BoolFromObject2 map[string]interface{}

	IntFromBool   bool
	IntFromBool2  bool
	IntFromFloat  float64
	IntFromString string

	FloatFromInt    int
	FloatFromBool   bool
	FloatFromString string
}

type DestinationCData struct {
	StringFromInt     String
	StringFromFloat   String
	StringFromBool    String
	StringFromObject  String
	StringFromObject2 String

	BoolFromInt     Bool
	BoolFromFloat   Bool
	BoolFromString  Bool
	BoolFromString2 Bool
	BoolFromString3 Bool
	BoolFromObject  Bool
	BoolFromObject2 Bool

	IntFromBool   Int
	IntFromBool2  Int
	IntFromFloat  Int
	IntFromString Int

	FloatFromInt    Float64
	FloatFromBool   Float64
	FloatFromString Float64
}

func TestCTypes(t *testing.T) {

	var src = SourceCData{
		StringFromInt:     2,
		StringFromFloat:   2.2,
		StringFromBool:    true,
		StringFromObject:  map[string]interface{}{},
		StringFromObject2: map[string]interface{}{"x": "x", "y": 1},

		BoolFromInt:     2,
		BoolFromFloat:   2.2,
		BoolFromString:  "2",
		BoolFromString2: "1",
		BoolFromString3: "true",
		BoolFromObject:  map[string]interface{}{},
		BoolFromObject2: map[string]interface{}{"x": "x", "y": 1},

		IntFromBool:   true,
		IntFromBool2:  false,
		IntFromFloat:  2.2,
		IntFromString: "2",

		FloatFromInt:    2,
		FloatFromBool:   true,
		FloatFromString: "2.2",
	}

	b, err := json.Marshal(src)
	if err != nil {
		t.Error(err)
	}

	dest := DestinationCData{}

	err = json.Unmarshal(b, &dest)
	if err != nil {
		t.Error(err)
	}

	// To string
	if dest.StringFromInt != "2" {
		t.Error("StringFromInt: " + string(dest.StringFromInt) + "/" + string(String("2")))
	}
	if dest.StringFromFloat != "2.2" {
		t.Error("StringFromFloat: " + string(dest.StringFromFloat) + "/" + string(String("2.2")))
	}
	if dest.StringFromBool != "true" {
		t.Error("StringFromBool: " + string(dest.StringFromBool) + "/" + string(String("true")))
	}
	if dest.StringFromObject != `` {
		t.Error("StringFromObject: " + string(dest.StringFromObject) + "/" + string(String(``)))
	}
	if dest.StringFromObject2 != `{"x":"x","y":1}` {
		t.Error("StringFromObject2: " + string(dest.StringFromObject2) + "/" + string(String(`{"x":"x","y":1}`)))
	}

	// To bool
	if dest.BoolFromInt != false {
		t.Error("BoolFromInt")
	}
	if dest.BoolFromFloat != false {
		t.Error("BoolFromFloat")
	}
	if dest.BoolFromString != false {
		t.Error("BoolFromString")
	}
	if dest.BoolFromString2 != true {
		t.Error("BoolFromString2")
	}
	if dest.BoolFromString3 != true {
		t.Error("BoolFromString3")
	}
	if dest.BoolFromObject != false {
		t.Error("BoolFromObject")
	}
	if dest.BoolFromObject2 != true {
		t.Error("BoolFromObject2")
	}

	// To int
	if dest.IntFromBool != 1 {
		t.Error("IntFromBool")
	}
	if dest.IntFromBool2 != 0 {
		t.Error("IntFromBool")
	}
	if dest.IntFromFloat != 2 {
		t.Error("IntFromFloat")
	}
	if dest.IntFromString != 2 {
		t.Error("IntFromString")
	}

	// To float
	if dest.FloatFromInt != 2 {
		t.Error("FloatFromInt")
	}
	if dest.FloatFromBool != 1 {
		t.Error("FloatFromBool")
	}
	if dest.FloatFromString != 2.2 {
		t.Error("FloatFromString")
	}
}
