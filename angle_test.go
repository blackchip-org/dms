package dms

import (
	"fmt"
	"testing"
)

func TestAngleString(t *testing.T) {
	tests := []struct {
		angle Angle
		str   string
	}{
		{Angle{Deg: "1"}, "1°"},
		{Angle{Sign: -1, Deg: "1"}, "-1°"},
		{Angle{Deg: "1", Min: "2"}, "1° 2′"},
		{Angle{Deg: "1", Min: "2", Sec: "3"}, "1° 2′ 3″"},
		{Angle{Deg: "1", Min: "2", Sec: "3", Hemi: "E"}, "1° 2′ 3″ E"},
		{Angle{Sign: -1, Deg: "1", Min: "2", Sec: "3", Hemi: "E"}, "1° 2′ 3″ E"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := test.angle.String()
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestAngleToDegrees(t *testing.T) {
	tests := []struct {
		angle Angle
		deg   string
	}{
		{Angle{Deg: "1"}, "1.000000"},
		{Angle{Deg: "1", Min: "3"}, "1.050000"},
		{Angle{Deg: "1", Min: "3", Sec: "6"}, "1.051667"},
		{Angle{Min: "3"}, "0.050000"},
		{Angle{Min: "3", Sec: "6"}, "0.051667"},
		{Angle{Sec: "6"}, "0.001667"},
		{Angle{Min: "3", Sec: "6", Hemi: "S"}, "-0.051667"},
		{Angle{Sign: -1, Min: "3", Sec: "6"}, "-0.051667"},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			deg := fmt.Sprintf("%.6f", test.angle.ToDegrees())
			if deg != test.deg {
				t.Errorf("\n have: %v \n want: %v", deg, test.deg)
			}
		})
	}
}

func TestAngleToMinutes(t *testing.T) {
	tests := []struct {
		angle Angle
		min   string
	}{
		{Angle{Deg: "1"}, "60.000"},
		{Angle{Deg: "1", Min: "3"}, "63.000"},
		{Angle{Deg: "1", Min: "3", Sec: "6"}, "63.100"},
		{Angle{Min: "3"}, "3.000"},
		{Angle{Min: "3", Sec: "6"}, "3.100"},
		{Angle{Sec: "6"}, "0.100"},
		{Angle{Min: "3", Sec: "6", Hemi: "S"}, "-3.100"},
		{Angle{Sign: -1, Min: "3", Sec: "6"}, "-3.100"},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			min := fmt.Sprintf("%.3f", test.angle.ToMinutes())
			if min != test.min {
				t.Errorf("\n have: %v \n want: %v", min, test.min)
			}
		})
	}
}

func TestAngleToSeconds(t *testing.T) {
	tests := []struct {
		angle Angle
		sec   string
	}{
		{Angle{Deg: "1"}, "3600.0"},
		{Angle{Deg: "1", Min: "3"}, "3780.0"},
		{Angle{Deg: "1", Min: "3", Sec: "6"}, "3786.0"},
		{Angle{Min: "3"}, "180.0"},
		{Angle{Min: "3", Sec: "6"}, "186.0"},
		{Angle{Sec: "6"}, "6.0"},
		{Angle{Min: "3", Sec: "6", Hemi: "S"}, "-186.0"},
		{Angle{Sign: -1, Min: "3", Sec: "6"}, "-186.0"},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			sec := fmt.Sprintf("%.1f", test.angle.ToSeconds())
			if sec != test.sec {
				t.Errorf("\n have: %v \n want: %v", sec, test.sec)
			}
		})
	}
}

func TestAngleToRadians(t *testing.T) {
	tests := []struct {
		angle Angle
		rad   string
	}{
		{Angle{Deg: "90"}, "1.570796"},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			rad := fmt.Sprintf("%.6f", test.angle.ToRadians())
			if rad != test.rad {
				t.Errorf("\n have: %v \n want: %v", rad, test.rad)
			}
		})
	}

}
