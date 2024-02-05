package dms

// import (
// 	"fmt"
// 	"testing"
// )

// func TestAngleToDegrees(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		deg   string
// 	}{
// 		{Parsed{Deg: "1"}, "1.000000"},
// 		{Parsed{Deg: "1", Min: "3"}, "1.050000"},
// 		{Parsed{Deg: "1", Min: "3", Sec: "6"}, "1.051667"},
// 		{Parsed{Min: "3"}, "0.050000"},
// 		{Parsed{Min: "3", Sec: "6"}, "0.051667"},
// 		{Parsed{Sec: "6"}, "0.001667"},
// 		{Parsed{Min: "3", Sec: "6", Hemi: "S"}, "-0.051667"},
// 		{Parsed{Sign: "-", Min: "3", Sec: "6"}, "-0.051667"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.deg, func(t *testing.T) {
// 			deg := fmt.Sprintf("%.6f", test.angle.ToDegrees())
// 			if deg != test.deg {
// 				t.Errorf("\n have: %v \n want: %v", deg, test.deg)
// 			}
// 		})
// 	}
// }

// func TestAngleToMinutes(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		min   string
// 	}{
// 		{Parsed{Deg: "1"}, "60.000"},
// 		{Parsed{Deg: "1", Min: "3"}, "63.000"},
// 		{Parsed{Deg: "1", Min: "3", Sec: "6"}, "63.100"},
// 		{Parsed{Min: "3"}, "3.000"},
// 		{Parsed{Min: "3", Sec: "6"}, "3.100"},
// 		{Parsed{Sec: "6"}, "0.100"},
// 		{Parsed{Min: "3", Sec: "6", Hemi: "S"}, "-3.100"},
// 		{Parsed{Sign: "-", Min: "3", Sec: "6"}, "-3.100"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.min, func(t *testing.T) {
// 			min := fmt.Sprintf("%.3f", test.angle.ToMinutes())
// 			if min != test.min {
// 				t.Errorf("\n have: %v \n want: %v", min, test.min)
// 			}
// 		})
// 	}
// }

// func TestAngleToSeconds(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		sec   string
// 	}{
// 		{Parsed{Deg: "1"}, "3600.0"},
// 		{Parsed{Deg: "1", Min: "3"}, "3780.0"},
// 		{Parsed{Deg: "1", Min: "3", Sec: "6"}, "3786.0"},
// 		{Parsed{Min: "3"}, "180.0"},
// 		{Parsed{Min: "3", Sec: "6"}, "186.0"},
// 		{Parsed{Sec: "6"}, "6.0"},
// 		{Parsed{Min: "3", Sec: "6", Hemi: "S"}, "-186.0"},
// 		{Parsed{Sign: "-", Min: "3", Sec: "6"}, "-186.0"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.sec, func(t *testing.T) {
// 			sec := fmt.Sprintf("%.1f", test.angle.ToSeconds())
// 			if sec != test.sec {
// 				t.Errorf("\n have: %v \n want: %v", sec, test.sec)
// 			}
// 		})
// 	}
// }

// func TestAngleToRadians(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		rad   string
// 	}{
// 		{Parsed{Deg: "90"}, "1.570796"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.rad, func(t *testing.T) {
// 			rad := fmt.Sprintf("%.6f", test.angle.ToRadians())
// 			if rad != test.rad {
// 				t.Errorf("\n have: %v \n want: %v", rad, test.rad)
// 			}
// 		})
// 	}
// }

// func TestAngleToFloatsErrors(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		err   string
// 	}{
// 		{Parsed{Deg: "x"}, "invalid degrees: x"},
// 		{Parsed{Sign: "+", Deg: "-1"}, "sign is positive but degrees are negative: -1"},
// 		{Parsed{Min: "x"}, "invalid minutes: x"},
// 		{Parsed{Min: "60"}, "invalid minutes: 60"},
// 		{Parsed{Sec: "x"}, "invalid seconds: x"},
// 		{Parsed{Sec: "60"}, "invalid seconds: 60"},
// 		{Parsed{Hemi: "x"}, "invalid hemisphere: x"},
// 		{Parsed{Sign: "-", Hemi: "N"}, "only one of '-' and 'N' allowed"},
// 		{Parsed{Sign: "+", Hemi: "S"}, "only one of '+' and 'S' allowed"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.err, func(t *testing.T) {
// 			_, _, _, _, err := test.angle.ToFloats()
// 			if err == nil {
// 				t.Fatal("expected error")
// 			}
// 			if err.Error() != test.err {
// 				t.Errorf("\n have: %v \n want: %v", err.Error(), test.err)
// 			}
// 		})
// 	}
// }

// func TestAsLat(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		lat   Parsed
// 	}{
// 		{Parsed{Deg: "1", Hemi: "X"}, Parsed{Deg: "1", Hemi: "X"}},
// 		{Parsed{Deg: "1"}, Parsed{Deg: "1", Hemi: "N"}},
// 		{Parsed{Sign: "+", Deg: "1"}, Parsed{Deg: "1", Hemi: "N"}},
// 		{Parsed{Sign: "-", Deg: "1"}, Parsed{Deg: "1", Hemi: "S"}},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("%+v", test.angle), func(t *testing.T) {
// 			lat := test.angle.AsLat()
// 			if lat != test.lat {
// 				t.Fatalf("\n have: %v \n want: %v", lat, test.lat)
// 			}
// 		})
// 	}
// }

// func TestAsLon(t *testing.T) {
// 	tests := []struct {
// 		angle Parsed
// 		lat   Parsed
// 	}{
// 		{Parsed{Deg: "1", Hemi: "X"}, Parsed{Deg: "1", Hemi: "X"}},
// 		{Parsed{Deg: "1"}, Parsed{Deg: "1", Hemi: "E"}},
// 		{Parsed{Sign: "+", Deg: "1"}, Parsed{Deg: "1", Hemi: "E"}},
// 		{Parsed{Sign: "-", Deg: "1"}, Parsed{Deg: "1", Hemi: "W"}},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("%+v", test.angle), func(t *testing.T) {
// 			lat := test.angle.AsLon()
// 			if lat != test.lat {
// 				t.Fatalf("\n have: %v \n want: %v", lat, test.lat)
// 			}
// 		})
// 	}
// }
