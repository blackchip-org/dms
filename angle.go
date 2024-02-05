package dms

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Angle struct {
	Sign int
	Deg  string
	Min  string
	Sec  string
	Hemi string
}

func (a Angle) AsLat() (Angle, error) {
	if a.Hemi != "" && a.Hemi != NorthType && a.Hemi != SouthType {
		return Angle{}, fmt.Errorf("invalid hemisphere: %v", a.Hemi)
	}
	if a.Sign < 0 {
		a.Hemi = SouthType
	} else if a.Sign >= 0 {
		a.Sign = 1
		a.Hemi = NorthType
	}
	return a, nil
}

func (a Angle) AsLon() (Angle, error) {
	if a.Hemi != "" && a.Hemi != EastType && a.Hemi != WestType {
		return Angle{}, fmt.Errorf("invalid hemisphere: %v", a.Hemi)
	}
	if a.Sign < 0 {
		a.Hemi = WestType
	} else if a.Sign >= 0 {
		a.Sign = 1
		a.Hemi = EastType
	}
	return a, nil
}

func (a Angle) ToFloats() (sign float64, deg float64, min float64, sec float64, err error) {
	if a.Deg != "" {
		deg, err = strconv.ParseFloat(a.Deg, 64)
		if err != nil {
			err = fmt.Errorf("invalid degrees: %v", a.Deg)
			return
		}
		if a.Sign == 1 && deg < 0 {
			err = fmt.Errorf("sign is positive but degrees are negative: %v", deg)
			return
		}
		if a.Sign != 0 {
			deg = math.Abs(deg)
		}
	}

	if a.Min != "" {
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

	sign = float64(a.Sign)
	if a.Hemi != "" {
		if a.Hemi != NorthType && a.Hemi != SouthType && a.Hemi != EastType && a.Hemi != WestType {
			err = fmt.Errorf("invalid hemisphere: %v", a.Hemi)
			return
		}
		if a.Sign == -1 && a.Hemi != SouthType && a.Hemi != WestType {
			err = fmt.Errorf("hemisphere mismatch: '-' and '%v'", a.Hemi)
		}
		if a.Sign == 1 && a.Hemi != NorthType && a.Hemi != EastType {
			err = fmt.Errorf("hemisphere mismatch: '+' and '%v'", a.Hemi)
		}
	}
	if sign == 0 {
		sign = 1
		if a.Hemi == SouthType || a.Hemi == WestType {
			sign = -1
		}
	}
	return
}

func (a Angle) String() string {
	var buf strings.Builder
	if a.Sign < 0 && a.Hemi == "" {
		buf.WriteRune('-')
	}
	buf.WriteString(a.Deg)
	buf.WriteRune('°')
	if a.Min != "" {
		buf.WriteRune(' ')
		buf.WriteString(a.Min)
		buf.WriteRune('′')
		if a.Sec != "" {
			buf.WriteRune(' ')
			buf.WriteString(a.Sec)
			buf.WriteRune('″')
		}
	}
	if a.Hemi != "" {
		buf.WriteRune(' ')
		buf.WriteString(a.Hemi)
	}
	return buf.String()
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
