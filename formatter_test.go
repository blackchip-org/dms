package dms

import "testing"

func TestFormat(t *testing.T) {
	var (
		def  = NewFormatter(SecType, 1)
		dms  = def.WithSymbols("d", "m", "")
		dm   = NewFormatter(MinType, 3)
		dd   = NewFormatter(DegType, 6)
		ddn  = NewFormatter(DegType, -1)
		mash = def.WithSep("")
	)

	tests := []struct {
		f      *Formatter
		angle  Angle
		result string
	}{
		{&def, NewAngle(1, 0, 0), "1° 0′ 0.0″"},
		{&def, NewAngle(1, 2, 0), "1° 2′ 0.0″"},
		{&def, NewAngle(1, 2, 3.33), "1° 2′ 3.3″"},
		{&def, NewAngle(-1, 2, 3.33), "-1° 2′ 3.3″"},
		{&def, NewAngle(1, -2, -3.33), "1° 2′ 3.3″"},
		{&def, NewAngle(-1, 2, 3.36), "-1° 2′ 3.4″"},
		{&def, NewAngle(1.051667, 0, 0), "1° 3′ 6.0″"},
		{&def, NewAngle(-1.051667, 0, 0), "-1° 3′ 6.0″"},
		{&def, NewAngle(1.5, 0, 0), "1° 30′ 0.0″"},
		{&def, NewAngle(1.5, 10, 0), "1° 40′ 0.0″"},
		{&def, NewAngle(-1.5, 10, 0), "-1° 40′ 0.0″"},
		{&def, NewAngle(-1.5, -10, 0), "-1° 40′ 0.0″"},
		{&def, NewAngle(1.5, 10.5, 10), "1° 40′ 40.0″"},
		{&def, NewAngle(0, 0, 75), "0° 1′ 15.0″"},
		{&def, NewAngle(0, 10, 135), "0° 12′ 15.0″"},
		{&def, NewAngle(0, 59, 135), "1° 1′ 15.0″"},

		{&dms, NewAngle(11, 22, 33.39), "11d 22m 33.4"},

		{&dm, NewAngle(1, 0, 0), "1° 0.000′"},
		{&dm, NewAngle(1, 2, 0), "1° 2.000′"},
		{&dm, NewAngle(1, 2, 6), "1° 2.100′"},

		{&dd, NewAngle(1, 0, 0), "1.000000°"},
		{&dd, NewAngle(1, 3, 0), "1.050000°"},
		{&dd, NewAngle(1, 3, 9), "1.052500°"},

		{&ddn, NewAngle(1, 0, 0), "1°"},
		{&mash, NewAngle(1, 2, 3.33), "1°2′3.3″"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			result := test.f.Format(test.angle)
			if result != test.result {
				t.Errorf("\n have: [%v] \n want: [%v]\n", result, test.result)
			}
		})
	}
}

func TestFormatLat(t *testing.T) {
	var (
		def  = NewFormatter(SecType, 1)
		dms  = def.WithSymbols("d", "m", "")
		dm   = NewFormatter(MinType, 3)
		dd   = NewFormatter(DegType, 6)
		ddn  = NewFormatter(DegType, -1)
		mash = def.WithSep("")
	)

	tests := []struct {
		f      *Formatter
		angle  Angle
		result string
	}{
		{&def, NewAngle(1, 0, 0), "1° 0′ 0.0″ N"},
		{&def, NewAngle(1, 2, 0), "1° 2′ 0.0″ N"},
		{&def, NewAngle(1, 2, 3.33), "1° 2′ 3.3″ N"},
		{&def, NewAngle(-1, 2, 3.33), "1° 2′ 3.3″ S"},
		{&def, NewAngle(1, -2, -3.33), "1° 2′ 3.3″ N"},
		{&def, NewAngle(-1, 2, 3.36), "1° 2′ 3.4″ S"},
		{&def, NewAngle(1.051667, 0, 0), "1° 3′ 6.0″ N"},
		{&def, NewAngle(-1.051667, 0, 0), "1° 3′ 6.0″ S"},
		{&def, NewAngle(1.5, 0, 0), "1° 30′ 0.0″ N"},
		{&def, NewAngle(1.5, 10, 0), "1° 40′ 0.0″ N"},
		{&def, NewAngle(-1.5, 10, 0), "1° 40′ 0.0″ S"},
		{&def, NewAngle(-1.5, -10, 0), "1° 40′ 0.0″ S"},
		{&def, NewAngle(1.5, 10.5, 10), "1° 40′ 40.0″ N"},
		{&def, NewAngle(0, 0, 75), "0° 1′ 15.0″ N"},
		{&def, NewAngle(0, 10, 135), "0° 12′ 15.0″ N"},
		{&def, NewAngle(0, 59, 135), "1° 1′ 15.0″ N"},

		{&dms, NewAngle(11, 22, 33.39), "11d 22m 33.4 N"},

		{&dm, NewAngle(1, 0, 0), "1° 0.000′ N"},
		{&dm, NewAngle(1, 2, 0), "1° 2.000′ N"},
		{&dm, NewAngle(1, 2, 6), "1° 2.100′ N"},

		{&dd, NewAngle(1, 0, 0), "1.000000° N"},
		{&dd, NewAngle(1, 3, 0), "1.050000° N"},
		{&dd, NewAngle(1, 3, 9), "1.052500° N"},

		{&ddn, NewAngle(1, 0, 0), "1° N"},

		{&mash, NewAngle(1, 2, 3.33), "1°2′3.3″N"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			result := test.f.FormatLat(test.angle)
			if result != test.result {
				t.Errorf("\n have: [%v] \n want: [%v]\n", result, test.result)
			}
		})
	}
}

func TestFormatLon(t *testing.T) {
	def := NewFormatter(SecType, 1)
	tests := []struct {
		f      *Formatter
		angle  Angle
		result string
	}{
		{&def, NewAngle(1, 0, 0), "1° 0′ 0.0″ E"},
		{&def, NewAngle(-1, 2, 0), "1° 2′ 0.0″ W"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			result := test.f.FormatLon(test.angle)
			if result != test.result {
				t.Errorf("\n have: [%v] \n want: [%v]\n", result, test.result)
			}
		})
	}
}
