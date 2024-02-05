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
		angle Parsed
		err   string
	}{
		{`1`, Parsed{Deg: "1"}, ""},
		{`+1`, Parsed{Hemi: "+", Deg: "1"}, ""},
		{`-1`, Parsed{Hemi: "-", Deg: "1"}, ""},
		{`1°`, Parsed{Deg: "1"}, ""},
		{`1d`, Parsed{Deg: "1"}, ""},
		{`1°N`, Parsed{Deg: "1", Hemi: "N"}, ""},
		{`1°S`, Parsed{Deg: "1", Hemi: "S"}, ""},
		{`1°E`, Parsed{Deg: "1", Hemi: "E"}, ""},
		{`1°W`, Parsed{Deg: "1", Hemi: "W"}, ""},
		{`1°2'`, Parsed{Deg: "1", Min: "2"}, ""},
		{`1°2′`, Parsed{Deg: "1", Min: "2"}, ""},
		{`1d2m`, Parsed{Deg: "1", Min: "2"}, ""},
		{`1°2'S`, Parsed{Deg: "1", Min: "2", Hemi: "S"}, ""},
		{`1°2'3"`, Parsed{Deg: "1", Min: "2", Sec: "3"}, ""},
		{`1°2′3″`, Parsed{Deg: "1", Min: "2", Sec: "3"}, ""},
		{`1°2′3″S`, Parsed{Deg: "1", Min: "2", Sec: "3", Hemi: "S"}, ""},
		{`1° 2′ 3″ S`, Parsed{Deg: "1", Min: "2", Sec: "3", Hemi: "S"}, ""},
		{`1.2`, Parsed{Deg: "1.2"}, ""},
		{`1.2°`, Parsed{Deg: "1.2"}, ""},
		{`1.2°S`, Parsed{Deg: "1.2", Hemi: "S"}, ""},
		{`1°2.3'`, Parsed{Deg: "1", Min: "2.3"}, ""},
		{`1°2.3'S`, Parsed{Deg: "1", Min: "2.3", Hemi: "S"}, ""},
		{`1°2'3.4"`, Parsed{Deg: "1", Min: "2", Sec: "3.4"}, ""},
		{`1°2'3.4"S`, Parsed{Deg: "1", Min: "2", Sec: "3.4", Hemi: "S"}, ""},
		{`9223372036854775807`, Parsed{Deg: "9223372036854775807"}, ""},

		{`x`, Parsed{}, `1:1: expected degree, got "x"`},
		{`+`, Parsed{}, `1:2: expected degree, got ""`},
		{`1x`, Parsed{}, `1:2: unexpected "x"`},
		{`1°x`, Parsed{}, `1:3: unexpected "x"`},
		{`1°2`, Parsed{}, `1:4: expected minute symbol, got ""`},
		{`1°2'3`, Parsed{}, `1:6: expected second symbol, got ""`},
		{`1.1x`, Parsed{}, `1:4: unexpected "x"`},
		{`1°2.2`, Parsed{}, `1:6: expected minute symbol, got ""`},
		{`1.1°2.2`, Parsed{}, `1:5: unexpected "2.2"`},
		{`1°2.2'3.3`, Parsed{}, `1:7: unexpected "3.3"`},
		{`9223372036854775808`, Parsed{}, `1:1: invalid degree "9223372036854775808"`},
		{bigReal, Parsed{}, `1:1: invalid degree "` + bigReal + `"`},
		{`1°60"`, Parsed{}, `1:3: invalid minute "60"`},
		{`1°59'60"`, Parsed{}, `1:6: invalid second "60"`},
		{`1°60.1"`, Parsed{}, `1:3: invalid minute "60.1"`},
		{`1°59'60.1"`, Parsed{}, `1:6: invalid second "60.1"`},
		{`-1°2'3.4"N`, Parsed{}, `1:10: only one of "-" or "N" are allowed`},
		{`+1°2'3.4"S`, Parsed{}, `1:10: only one of "+" or "S" are allowed`},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			p := NewDefaultParser()
			angle, err := p.Parse(test.input)
			errMsg := ""
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != test.err {
				t.Fatalf("\n have err: %v \n want err: %v", errMsg, test.err)
			}
			if angle != test.angle {
				t.Errorf("\n have: %v \n want: %v", angle, test.angle)
			}
		})
	}
}

func ExampleParser_Parse() {
	p := NewDefaultParser()
	a, err := p.ParseAngle("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
}
