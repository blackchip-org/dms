package dms

import "testing"

func TestAngleAdd(t *testing.T) {
	tests := []struct {
		a   Angle
		b   Angle
		str string
	}{
		{NewAngle(1, 2, 3), NewAngle(4, 5, 6), "5° 7′ 9.0″ N"},
		{NewAngle(-1, 0, 0), NewAngle(1, 0, 0), "0° 0′ 0.0″ N"},
		{NewAngle(-1, 15, 0), NewAngle(1, 15, 0), "0° 0′ 0.0″ N"},
	}

	f := NewFormatter(SecType, 1)
	for _, test := range tests {
		c := test.a.Add(test.b)
		str := f.FormatLat(c)
		if str != test.str {
			t.Errorf("\n have: %v \n want: %v", str, test.str)
		}
	}
}
