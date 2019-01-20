package main

import (
	"encoding/json"
	"github.com/Jleagle/unmarshal-go/unmarshal"
	"testing"
)

type SourceData struct {
	StringFromInt   int     `json:"string_from_int"`
	StringFromFloat float64 `json:"string_from_float"`
	StringFromBool  bool    `json:"string_from_bool"`

	BoolFromInt    int
	BoolFromFloat  float64
	BoolFromString string

	IntFromBool   bool
	IntFromFloat  float64
	IntFromString string

	FloatFromInt    int
	FloatFromBool   bool
	FloatFromString string
}

type DestinationData struct {
	StringFromInt   string `json:"string_from_int"`
	StringFromFloat string `json:"string_from_float"`
	StringFromBool  string `json:"string_from_bool"`

	BoolFromInt    bool
	BoolFromFloat  bool
	BoolFromString bool

	IntFromBool   int
	IntFromFloat  int
	IntFromString int

	FloatFromInt    float64
	FloatFromBool   float64
	FloatFromString float64
}

func Test(t *testing.T) {

	var src = SourceData{
		StringFromInt:   2,
		StringFromFloat: 2.2,
		StringFromBool:  true,
		BoolFromInt:     2,
		BoolFromFloat:   2.2,
		BoolFromString:  "2",
		IntFromBool:     true,
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

	dest := DestinationData{}

	err = unmarshal.Unmarshal(b, &dest)
	if err != nil {
		t.Error(err)
	}
}

func TestMap(t *testing.T) {

	var src = SourceData{
		StringFromInt:   2,
		StringFromFloat: 2.2,
		StringFromBool:  true,
		BoolFromInt:     2,
		BoolFromFloat:   2.2,
		BoolFromString:  "2",
		IntFromBool:     true,
		IntFromFloat:    2.2,
		IntFromString:   "2",
		FloatFromInt:    2,
		FloatFromBool:   true,
		FloatFromString: "2.2",
	}

	srcMap := map[string]SourceData{
		"1": src,
		"2": src,
	}

	b, err := json.Marshal(srcMap)
	if err != nil {
		t.Error(err)
	}

	destMap := map[string]DestinationData{}

	err = unmarshal.Unmarshal(b, &destMap)
	if err != nil {
		t.Error(err)
	}

	if destMap["1"].StringFromInt != "2" {
		t.Error("not 2")
	}
}
