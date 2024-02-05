package dms

import (
	"fmt"
	"math"
	"strconv"
)

type Angle struct {
	Sign string
	Deg  string
	Min  string
	Sec  string
	Hemi string
}

func NewAngle(sign string, deg string, min string, sec string, hemi string) Angle {
	return Angle{
		Sign: sign,
		Deg:  deg,
		Min:  min,
		Sec:  sec,
		Hemi: hemi,
	}
}

func (a Angle) AsLat() Angle {
	if a.Hemi != "" {
		return a
	}
	if a.Sign == "-" {
		a.Hemi = SouthType
	} else {
		a.Hemi = NorthType
	}
	a.Sign = ""
	return a
}

func (a Angle) AsLon() Angle {
	if a.Hemi != "" {
		return a
	}
	if a.Sign == "-" {
		a.Hemi = WestType
	} else {
		a.Hemi = EastType
	}
	a.Sign = ""
	return a
}

func (a Angle) ToFloats() (sign float64, deg float64, min float64, sec float64, err error) {
	switch a.Sign {
	case "-":
		sign = -1
	case "+":
		sign = 1
	case "":
		sign = 0
	default:
		err = fmt.Errorf("invalid sign: %v", a.Sign)
		return
	}

	if a.Deg != "" {
		deg, err = strconv.ParseFloat(a.Deg, 64)
		if err != nil {
			err = fmt.Errorf("invalid degrees: %v", a.Deg)
			return
		}
		if a.Sign == "+" && deg < 0 {
			err = fmt.Errorf("sign is positive but degrees are negative: %v", deg)
			return
		}
		if a.Sign != "" {
			deg = math.Abs(deg)
		}
	}

	if a.Min != "" {
		if deg != math.Trunc(deg) {
			err = fmt.Errorf("decimal degrees %v with minutes %v", a.Deg, a.Min)
			return
		}
		min, err = strconv.ParseFloat(a.Min, 64)
		if err != nil {
			err = fmt.Errorf("invalid minutes: %v", a.Min)
			return
		}
		if min < 0 || min >= 60 {
			err = fmt.Errorf("invalid minutes: %v", a.Min)
			return
		}
	}

	if a.Sec != "" {
		if min != math.Trunc(min) {
			err = fmt.Errorf("decimal minutes %v with seconds %v", a.Min, a.Sec)
			return
		}
		sec, err = strconv.ParseFloat(a.Sec, 64)
		if err != nil {
			err = fmt.Errorf("invalid seconds: %v", a.Sec)
			return
		}
		sec = math.Abs(sec)
		if sec < 0 || sec >= 60 {
			err = fmt.Errorf("invalid seconds: %v", a.Sec)
		}
	}

	if a.Hemi != "" {
		if a.Hemi != NorthType && a.Hemi != SouthType && a.Hemi != EastType && a.Hemi != WestType {
			err = fmt.Errorf("invalid hemisphere: %v", a.Hemi)
			return
		}
		if a.Sign != "" {
			err = fmt.Errorf("only one of '%v' and '%v' allowed", a.Sign, a.Hemi)
			return
		}
	}
	if sign == 0 {
		sign = 1
		if a.Hemi == SouthType || a.Hemi == WestType {
			sign = -1
		}
	}

	deg, min, sec = normalizeFloats(deg, min, sec)
	return
}

func normalizeFloats(deg, min, sec float64) (float64, float64, float64) {
	if ideg, fdeg := math.Modf(deg); fdeg != 0 && min == 0 && sec == 0 {
		deg = ideg
		min = math.Abs(fdeg) * 60
	}
	if imin, fmin := math.Modf(min); fmin != 0 && sec == 0 {
		min = imin
		sec = math.Abs(fmin) * 60
	}
	return deg, min, sec
}

func (a Angle) ToDegrees() float64 {
	sign, deg, min, sec, err := a.ToFloats()
	if err != nil {
		panic(err)
	}
	return sign * (deg + (min / 60) + (sec / 3600))
}

func (a Angle) ToMinutes() float64 {
	return a.ToDegrees() * 60
}

func (a Angle) ToSeconds() float64 {
	return a.ToDegrees() * 3600
}

func (a Angle) ToRadians() float64 {
	return a.ToDegrees() * (math.Pi / 180)
}
