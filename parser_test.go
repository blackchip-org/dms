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
		angle Angle
		err   string
	}{
		{`1`, Angle{Deg: "1"}, ""},
		{`+1`, Angle{Sign: "+", Deg: "1"}, ""},
		{`-1`, Angle{Sign: "-", Deg: "1"}, ""},
		{`1°`, Angle{Deg: "1"}, ""},
		{`1d`, Angle{Deg: "1"}, ""},
		{`1°N`, Angle{Deg: "1", Hemi: "N"}, ""},
		{`1°S`, Angle{Deg: "1", Hemi: "S"}, ""},
		{`1°E`, Angle{Deg: "1", Hemi: "E"}, ""},
		{`1°W`, Angle{Deg: "1", Hemi: "W"}, ""},
		{`1°2'`, Angle{Deg: "1", Min: "2"}, ""},
		{`1°2′`, Angle{Deg: "1", Min: "2"}, ""},
		{`1d2m`, Angle{Deg: "1", Min: "2"}, ""},
		{`1°2'S`, Angle{Deg: "1", Min: "2", Hemi: "S"}, ""},
		{`1°2'3"`, Angle{Deg: "1", Min: "2", Sec: "3"}, ""},
		{`1°2′3″`, Angle{Deg: "1", Min: "2", Sec: "3"}, ""},
		{`1°2′3″S`, Angle{Deg: "1", Min: "2", Sec: "3", Hemi: "S"}, ""},
		{`1° 2′ 3″ S`, Angle{Deg: "1", Min: "2", Sec: "3", Hemi: "S"}, ""},
		{`1.2`, Angle{Deg: "1.2"}, ""},
		{`1.2°`, Angle{Deg: "1.2"}, ""},
		{`1.2°S`, Angle{Deg: "1.2", Hemi: "S"}, ""},
		{`1°2.3'`, Angle{Deg: "1", Min: "2.3"}, ""},
		{`1°2.3'S`, Angle{Deg: "1", Min: "2.3", Hemi: "S"}, ""},
		{`1°2'3.4"`, Angle{Deg: "1", Min: "2", Sec: "3.4"}, ""},
		{`1°2'3.4"S`, Angle{Deg: "1", Min: "2", Sec: "3.4", Hemi: "S"}, ""},
		{`9223372036854775807`, Angle{Deg: "9223372036854775807"}, ""},

		{`x`, Angle{}, `1:1: expected degree, got "x"`},
		{`+`, Angle{}, `1:2: expected degree, got ""`},
		{`1x`, Angle{}, `1:2: unexpected "x"`},
		{`1°x`, Angle{}, `1:3: unexpected "x"`},
		{`1°2`, Angle{}, `1:4: expected minute symbol, got ""`},
		{`1°2'3`, Angle{}, `1:6: expected second symbol, got ""`},
		{`1.1x`, Angle{}, `1:4: unexpected "x"`},
		{`1°2.2`, Angle{}, `1:6: expected minute symbol, got ""`},
		{`1.1°2.2`, Angle{}, `1:5: unexpected "2.2"`},
		{`1°2.2'3.3`, Angle{}, `1:7: unexpected "3.3"`},
		{`9223372036854775808`, Angle{}, `1:1: invalid degree "9223372036854775808"`},
		{bigReal, Angle{}, `1:1: invalid degree "` + bigReal + `"`},
		{`1°60"`, Angle{}, `1:3: invalid minute "60"`},
		{`1°59'60"`, Angle{}, `1:6: invalid second "60"`},
		{`1°60.1"`, Angle{}, `1:3: invalid minute "60.1"`},
		{`1°59'60.1"`, Angle{}, `1:6: invalid second "60.1"`},
		{`-1°2'3.4"N`, Angle{}, `1:1: only one of "-" or "N" allowed`},
		{`+1°2'3.4"S`, Angle{}, `1:1: only one of "+" or "S" allowed`},
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
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
}
