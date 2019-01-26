package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SourceCData struct {
	StringFromInt   int     `json:"string_from_int"`
	StringFromFloat float64 `json:"string_from_float"`
	StringFromBool  bool    `json:"string_from_bool"`

	// BoolFromInt    int
	// BoolFromFloat  float64
	// BoolFromString string
	//
	// IntFromBool   bool
	// IntFromFloat  float64
	// IntFromString string
	//
	// FloatFromInt    int
	// FloatFromBool   bool
	// FloatFromString string
}

type DestinationCData struct {
	StringFromInt   CString `json:"string_from_int"`
	StringFromFloat CString `json:"string_from_float"`
	StringFromBool  CString `json:"string_from_bool"`

	// BoolFromInt    CBool
	// BoolFromFloat  CBool
	// BoolFromString CBool
	//
	// IntFromBool   CInt
	// IntFromFloat  CInt
	// IntFromString CInt
	//
	// FloatFromInt    CFloat
	// FloatFromBool   CFloat
	// FloatFromString CFloat
}

func TestCTypes(t *testing.T) {

	var src = SourceCData{
		StringFromInt:   2,
		StringFromFloat: 2.2,
		StringFromBool:  true,
		// BoolFromInt:     2,
		// BoolFromFloat:   2.2,
		// BoolFromString:  "2",
		// IntFromBool:     true,
		// IntFromFloat:    2.2,
		// IntFromString:   "2",
		// FloatFromInt:    2,
		// FloatFromBool:   true,
		// FloatFromString: "2.2",
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

	fmt.Println(dest)

	if dest.StringFromInt != CString("2") {
		t.Error("StringFromInt: " + string(dest.StringFromInt) + "/" + string(CString("2")))
	}
	if dest.StringFromFloat != CString("2.2") {
		t.Error("StringFromFloat: " + string(dest.StringFromFloat) + "/" + string(CString("2.2")))
	}
	if dest.StringFromBool != CString("true") {
		t.Error("StringFromBool: " + string(dest.StringFromBool) + "/" + string(CString("true")))
	}
}
