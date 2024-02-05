package dms

import (
	"fmt"
	"testing"
)

func TestFormatter(t *testing.T) {
	def := NewFormatter(SecType, 1)
	dms := def.WithSymbols("d", "m", "s")

	dm := NewFormatter(MinType, 3)
	dd := NewFormatter(DegType, 6)

	sign := def.WithSign(true)
	mash := def.WithSep("")

	tests := []struct {
		f      *Formatter
		angle  Angle
		result string
	}{
		{&def, Angle{Deg: "1"}, "1° 0′ 0.0″"},
		{&def, Angle{Deg: "1", Min: "2"}, "1° 2′ 0.0″"},
		{&def, Angle{Deg: "1", Min: "2", Sec: "3.39"}, "1° 2′ 3.4″"},
		{&def, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "N"}, "1° 2′ 3.4″ N"},
		{&def, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "S"}, "1° 2′ 3.4″ S"},
		{&def, Angle{Deg: "-1", Min: "2", Sec: "3.39"}, "-1° 2′ 3.4″"},
		{&def, Angle{Deg: "1.051667", Hemi: "S"}, "1° 3′ 6.0″ S"},
		{&def, Angle{Deg: "12.582439", Hemi: "W"}, "12° 34′ 56.8″ W"},
		{&def, Angle{Deg: "12", Min: "34.56", Hemi: "W"}, "12° 34′ 33.6″ W"},

		{&dms, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "S"}, "1d 2m 3.4s S"},

		{&dm, Angle{Deg: "1"}, "1° 0.000′"},
		{&dm, Angle{Deg: "1", Min: "2"}, "1° 2.000′"},
		{&dm, Angle{Deg: "1", Min: "2", Sec: "6"}, "1° 2.100′"},

		{&dd, Angle{Deg: "1"}, "1.000000°"},
		{&dd, Angle{Deg: "1", Min: "3"}, "1.050000°"},
		{&dd, Angle{Deg: "1", Min: "3", Sec: "6"}, "1.051667°"},

		{&sign, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "N"}, "1° 2′ 3.4″"},
		{&sign, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "S"}, "-1° 2′ 3.4″"},
		{&mash, Angle{Deg: "1", Min: "2", Sec: "3.39", Hemi: "S"}, "1°2′3.4″S"},
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

func TestFormatterPanic(t *testing.T) {
	tests := []struct {
		angle Angle
		msg   string
	}{
		{Angle{Deg: "12.34", Min: "56.78", Hemi: "W"}, "decimal degrees 12.34 with minutes 56.78"},
		{Angle{Deg: "12", Min: "56.78", Sec: "99.99", Hemi: "W"}, "decimal minutes 56.78 with seconds 99.99"},
		{Angle{Deg: "x"}, "invalid degrees: x"},
		{Angle{Sign: "x", Deg: "1"}, "invalid sign: x"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					msg := fmt.Sprintf("%v", r)
					if msg != test.msg {
						t.Errorf("\n have: %v \n want: %v", msg, test.msg)
					}
				}
			}()
			NewFormatter(SecType, 1).Format(test.angle)
			t.Error("did not panic")
		})
	}
}

func ExampleFormatter_Format() {
	f := NewFormatter(SecType, 1)
	fmt.Println(f.Format(Angle{Deg: "1.051667", Hemi: "S"}))

	// Output:
	// 1° 3′ 6.0″ S
}

func TestFormatLat(t *testing.T) {
	fsec := NewFormatter(SecType, 1)
	fdeg := NewFormatter(DegType, 6)

	tests := []struct {
		f   *Formatter
		deg float64
		min float64
		sec float64
		str string
	}{
		{&fsec, 1, 0, 0, "1° 0′ 0.0″ N"},
		{&fsec, 1.5, 0, 0, "1° 30′ 0.0″ N"},
		{&fsec, 1, 2, 0, "1° 2′ 0.0″ N"},
		{&fsec, 1, 2, 3, "1° 2′ 3.0″ N"},
		{&fsec, -1, 0, 0, "1° 0′ 0.0″ S"},
		{&fsec, -1, 2, 0, "1° 2′ 0.0″ S"},
		{&fsec, -1, 2, 3, "1° 2′ 3.0″ S"},

		{&fdeg, 1, 0, 0, "1.000000° N"},
		{&fdeg, 1, 3, 0, "1.050000° N"},
		{&fdeg, 1, 3, 6, "1.051667° N"},
		{&fdeg, -1, 0, 0, "1.000000° S"},
		{&fdeg, -1, 3, 0, "1.050000° S"},
		{&fdeg, -1, -3, -6, "1.051667° S"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := test.f.FormatLat(test.deg, test.min, test.sec)
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestFormatLon(t *testing.T) {
	fsec := NewFormatter(SecType, 1)
	fdeg := NewFormatter(DegType, 6)

	tests := []struct {
		f   *Formatter
		deg float64
		min float64
		sec float64
		str string
	}{
		{&fsec, 1, 0, 0, "1° 0′ 0.0″ E"},
		{&fsec, 1, 2, 0, "1° 2′ 0.0″ E"},
		{&fsec, 1, 2, 3, "1° 2′ 3.0″ E"},
		{&fsec, -1, 0, 0, "1° 0′ 0.0″ W"},
		{&fsec, -1, 2, 0, "1° 2′ 0.0″ W"},
		{&fsec, -1, 2, 3, "1° 2′ 3.0″ W"},

		{&fdeg, 1, 0, 0, "1.000000° E"},
		{&fdeg, 1, 3, 0, "1.050000° E"},
		{&fdeg, 1, 3, 6, "1.051667° E"},
		{&fdeg, -1, 0, 0, "1.000000° W"},
		{&fdeg, -1, 3, 0, "1.050000° W"},
		{&fdeg, -1, -3, -6, "1.051667° W"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := test.f.FormatLon(test.deg, test.min, test.sec)
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}
