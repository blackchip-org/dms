package dms

import (
	"fmt"
	"strings"
	"testing"
)

var bigReal = "1" + strings.Repeat("0", 309) + ".1"

func TestParser(t *testing.T) {
	tests := []struct {
		input string
		angle Fields
		err   string
	}{
		{`1`, Fields{Deg: "1"}, ""},
		{`+1`, Fields{Hemi: "+", Deg: "1"}, ""},
		{`-1`, Fields{Hemi: "-", Deg: "1"}, ""},
		{`1°`, Fields{Deg: "1", DegSym: "°"}, ""},
		{`1d`, Fields{Deg: "1", DegSym: "d"}, ""},
		{`1°N`, Fields{Deg: "1", DegSym: "°", Hemi: "N"}, ""},
		{`1°S`, Fields{Deg: "1", DegSym: "°", Hemi: "S"}, ""},
		{`1°E`, Fields{Deg: "1", DegSym: "°", Hemi: "E"}, ""},
		{`1°W`, Fields{Deg: "1", DegSym: "°", Hemi: "W"}, ""},
		{`1°2'`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "'"}, ""},
		{`1°2′`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "′"}, ""},
		{`1d2m`, Fields{Deg: "1", DegSym: "d", Min: "2", MinSym: "m"}, ""},
		{`1°2'S`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "'", Hemi: "S"}, ""},
		{`1°2'3"`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "'", Sec: "3", SecSym: `"`}, ""},
		{`1°2′3″`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "′", Sec: "3", SecSym: "″"}, ""},
		{`1°2′3″S`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "′", Sec: "3", SecSym: "″", Hemi: "S"}, ""},
		{`1° 2′ 3″ S`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "′", Sec: "3", SecSym: "″", Hemi: "S"}, ""},
		{`1.2`, Fields{Deg: "1.2"}, ""},
		{`1.2°`, Fields{Deg: "1.2", DegSym: "°"}, ""},
		{`1.2°S`, Fields{Deg: "1.2", DegSym: "°", Hemi: "S"}, ""},
		{`1°2.3'`, Fields{Deg: "1", DegSym: "°", Min: "2.3", MinSym: "'"}, ""},
		{`1°2.3'S`, Fields{Deg: "1", DegSym: "°", Min: "2.3", MinSym: "'", Hemi: "S"}, ""},
		{`1°2'3.4"`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "'", Sec: "3.4", SecSym: `"`}, ""},
		{`1°2'3.4"S`, Fields{Deg: "1", DegSym: "°", Min: "2", MinSym: "'", Sec: "3.4", SecSym: `"`, Hemi: "S"}, ""},
		{`9223372036854775807`, Fields{Deg: "9223372036854775807"}, ""},

		{`x`, Fields{}, `1:1: expected degree, got "x"`},
		{`+`, Fields{}, `1:2: expected degree, got ""`},
		{`1x`, Fields{}, `1:2: unexpected "x"`},
		{`1°x`, Fields{}, `1:3: unexpected "x"`},
		{`1°2`, Fields{}, `1:4: expected minute symbol, got ""`},
		{`1°2'3`, Fields{}, `1:6: expected second symbol, got ""`},
		{`1.1x`, Fields{}, `1:4: unexpected "x"`},
		{`1°2.2`, Fields{}, `1:6: expected minute symbol, got ""`},
		{`1.1°2.2`, Fields{}, `1:5: unexpected "2.2"`},
		{`1°2.2'3.3`, Fields{}, `1:7: unexpected "3.3"`},
		{`9223372036854775808`, Fields{}, `1:1: invalid degree "9223372036854775808"`},
		{bigReal, Fields{}, `1:1: invalid degree "` + bigReal + `"`},
		{`1°60"`, Fields{}, `1:3: invalid minute "60"`},
		{`1°59'60"`, Fields{}, `1:6: invalid second "60"`},
		{`1°60.1"`, Fields{}, `1:3: invalid minute "60.1"`},
		{`1°59'60.1"`, Fields{}, `1:6: invalid second "60.1"`},
		{`-1°2'3.4"N`, Fields{}, `1:10: only one of "-" or "N" are allowed`},
		{`+1°2'3.4"S`, Fields{}, `1:10: only one of "+" or "S" are allowed`},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			p := NewDefaultParser()
			angle, err := p.ParseFields(test.input)
			errMsg := ""
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != test.err {
				t.Fatalf("\n have err: %v \n want err: %v", errMsg, test.err)
			}
			if angle != test.angle {
				t.Errorf("\n have: %+v \n want: %+v", angle, test.angle)
			}
		})
	}
}

func ExampleParser_Parse() {
	p := NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.Degrees())

	// Output:
	// -1.051667
}
