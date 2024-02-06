package dms

import (
	"fmt"
	"testing"
)

func TestDegrees(t *testing.T) {
	tests := []struct {
		angle Angle
		str   string
	}{
		{NewAngle(1, 0, 0), "1.000000"},
		{NewAngle(1, 3, 0), "1.050000"},
		{NewAngle(1, 3, 9), "1.052500"},
		{NewAngle(0, 3, 0), "0.050000"},
		{NewAngle(0, 3, 9), "0.052500"},
		{NewAngle(0, 0, 9), "0.002500"},
		{NewAngle(-1, 3, 9), "-1.052500"},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			str := fmt.Sprintf("%.6f", test.angle.Degrees())
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestMinutes(t *testing.T) {
	tests := []struct {
		angle Angle
		str   string
	}{
		{NewAngle(1, 0, 0), "60.000"},
		{NewAngle(1, 3, 0), "63.000"},
		{NewAngle(1, 3, 9), "63.150"},
		{NewAngle(0, 3, 0), "3.000"},
		{NewAngle(0, 3, 9), "3.150"},
		{NewAngle(0, 0, 9), "0.150"},
		{NewAngle(-1, 3, 9), "-63.150"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := fmt.Sprintf("%.3f", test.angle.Minutes())
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestSeconds(t *testing.T) {
	tests := []struct {
		angle Angle
		str   string
	}{
		{NewAngle(1, 0, 0), "3600.0"},
		{NewAngle(1, 3, 0), "3780.0"},
		{NewAngle(1, 3, 9), "3789.0"},
		{NewAngle(0, 3, 0), "180.0"},
		{NewAngle(0, 3, 9), "189.0"},
		{NewAngle(0, 0, 9), "9.0"},
		{NewAngle(-1, 3, 9), "-3789.0"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := fmt.Sprintf("%.1f", test.angle.Seconds())
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

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
