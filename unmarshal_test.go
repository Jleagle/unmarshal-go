package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(dest)
}
