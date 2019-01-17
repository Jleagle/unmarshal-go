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
	src.StringFromInt = 2
	src.StringFromFloat = 2.2
	src.StringFromBool = true
	src.BoolFromInt = 2
	src.BoolFromFloat = 2.2
	src.BoolFromString = "2"
	src.IntFromBool = true
	src.IntFromFloat = 2.2
	src.IntFromString = "2"
	src.FloatFromInt = 2
	src.FloatFromBool = true
	src.FloatFromString = "2.2"

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
