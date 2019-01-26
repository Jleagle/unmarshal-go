package main

import (
	"encoding/json"
	"testing"
)

type SourceCData struct {
	StringFromInt   int
	StringFromFloat float64
	StringFromBool  bool

	BoolFromInt     int
	BoolFromFloat   float64
	BoolFromString  string
	BoolFromString2 string
	BoolFromString3 string

	IntFromBool   bool
	IntFromBool2  bool
	IntFromFloat  float64
	IntFromString string

	FloatFromInt    int
	FloatFromBool   bool
	FloatFromString string
}

type DestinationCData struct {
	StringFromInt   CString
	StringFromFloat CString
	StringFromBool  CString

	BoolFromInt     CBool
	BoolFromFloat   CBool
	BoolFromString  CBool
	BoolFromString2 CBool
	BoolFromString3 CBool

	IntFromBool   CInt
	IntFromBool2  CInt
	IntFromFloat  CInt
	IntFromString CInt

	FloatFromInt    CFloat
	FloatFromBool   CFloat
	FloatFromString CFloat
}

func TestCTypes(t *testing.T) {

	var src = SourceCData{
		StringFromInt:   2,
		StringFromFloat: 2.2,
		StringFromBool:  true,
		BoolFromInt:     2,
		BoolFromFloat:   2.2,
		BoolFromString:  "2",
		BoolFromString2: "1",
		BoolFromString3: "true",
		IntFromBool:     true,
		IntFromBool2:    false,
		IntFromFloat:    2.2,
		IntFromString:   "2",
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

	if dest.StringFromInt != "2" {
		t.Error("StringFromInt: " + string(dest.StringFromInt) + "/" + string(CString("2")))
	}
	if dest.StringFromFloat != "2.2" {
		t.Error("StringFromFloat: " + string(dest.StringFromFloat) + "/" + string(CString("2.2")))
	}
	if dest.StringFromBool != "true" {
		t.Error("StringFromBool: " + string(dest.StringFromBool) + "/" + string(CString("true")))
	}

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
		t.Error("BoolFromString")
	}
	if dest.BoolFromString3 != true {
		t.Error("BoolFromString")
	}

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
