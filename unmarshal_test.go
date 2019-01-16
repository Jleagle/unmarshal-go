package main

import (
	"encoding/json"
	"github.com/Jleagle/unmarshal-go/unmarshal"
	"testing"
)

type DestinationData struct {
	StringFromInt   string `json:"string_from_int"`
	StringFromFloat string
	StringFromBool  string

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

type SourceData struct {
	StringFromInt   int
	StringFromFloat float64
	StringFromBool  bool

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

func Test(t *testing.T) {

	src := SourceData{}
	src.BoolFromInt = 1
	src.BoolFromFloat = 1.5
	src.BoolFromString = "1"

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
