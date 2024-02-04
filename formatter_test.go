package dms

import "testing"

func TestFormatter(t *testing.T) {
	def := NewFormatter(SecType, 1)

	tests := []struct {
		f      *Formatter
		angle  Angle
		result string
	}{
		{&def, Angle{Deg: "1"}, "1° 0′ 0.0″"},
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
