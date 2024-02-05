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

func TestAngleToFloatsErrors(t *testing.T) {
	tests := []struct {
		angle Angle
		err   string
	}{
		{Angle{Deg: "x"}, "invalid degrees: x"},
		{Angle{Sign: 1, Deg: "-1"}, "sign is positive but degrees are negative: -1"},
		{Angle{Min: "x"}, "invalid minutes: x"},
		{Angle{Min: "60"}, "invalid minutes: 60"},
		{Angle{Sec: "x"}, "invalid seconds: x"},
		{Angle{Sec: "60"}, "invalid seconds: 60"},
		{Angle{Hemi: "x"}, "invalid hemisphere: x"},
		{Angle{Sign: -1, Hemi: "N"}, "hemisphere mismatch: '-' and 'N'"},
		{Angle{Sign: 1, Hemi: "S"}, "hemisphere mismatch: '+' and 'S'"},
	}

	for _, test := range tests {
		t.Run(test.err, func(t *testing.T) {
			_, _, _, _, err := test.angle.ToFloats()
			if err == nil {
				t.Fatal("expected error")
			}
			if err.Error() != test.err {
				t.Errorf("\n have: %v \n want: %v", err.Error(), test.err)
			}
		})
	}
}

func TestAsLat(t *testing.T) {
	tests := []struct {
		angle Angle
		lat   Angle
		err   string
	}{
		{Angle{Deg: "1", Hemi: "W"}, Angle{}, "invalid hemisphere: W"},
		{Angle{Deg: "1", Hemi: "E"}, Angle{}, "invalid hemisphere: E"},
		{Angle{Deg: "1", Hemi: "X"}, Angle{}, "invalid hemisphere: X"},
		{Angle{Deg: "1"}, Angle{Sign: 1, Deg: "1", Hemi: "N"}, ""},
		{Angle{Sign: 1, Deg: "1"}, Angle{Sign: 1, Deg: "1", Hemi: "N"}, ""},
		{Angle{Sign: -1, Deg: "1"}, Angle{Sign: -1, Deg: "1", Hemi: "S"}, ""},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			lat, err := test.angle.AsLat()
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != test.err {
				t.Fatalf("\n have err: %v \n want err: %v", errMsg, test.err)
			}
			if lat != test.lat {
				t.Fatalf("\n have: %v \n want: %v", lat, test.lat)
			}
		})
	}
}

func TestAsLon(t *testing.T) {
	tests := []struct {
		angle Angle
		lat   Angle
		err   string
	}{
		{Angle{Deg: "1", Hemi: "N"}, Angle{}, "invalid hemisphere: N"},
		{Angle{Deg: "1", Hemi: "S"}, Angle{}, "invalid hemisphere: S"},
		{Angle{Deg: "1", Hemi: "X"}, Angle{}, "invalid hemisphere: X"},
		{Angle{Deg: "1"}, Angle{Sign: 1, Deg: "1", Hemi: "E"}, ""},
		{Angle{Sign: 1, Deg: "1"}, Angle{Sign: 1, Deg: "1", Hemi: "E"}, ""},
		{Angle{Sign: -1, Deg: "1"}, Angle{Sign: -1, Deg: "1", Hemi: "W"}, ""},
	}

	for _, test := range tests {
		t.Run(test.angle.String(), func(t *testing.T) {
			lat, err := test.angle.AsLon()
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != test.err {
				t.Fatalf("\n have err: %v \n want err: %v", errMsg, test.err)
			}
			if lat != test.lat {
				t.Fatalf("\n have: %v \n want: %v", lat, test.lat)
			}
		})
	}
}
