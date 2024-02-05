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
	defer func() {
		if r := recover(); r != nil {
			// good!
		}
	}()
	NewFormatter(SecType, 1).Format(Angle{Deg: "x"})
	t.Error("did not panic")
}

func ExampleFormatter_Format() {
	f := NewFormatter(SecType, 1)
	fmt.Println(f.Format(Angle{Deg: "1.051667", Hemi: "S"}))

	// Output:
	// 1° 3′ 6.0″ S
}
