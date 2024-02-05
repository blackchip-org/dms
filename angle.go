package dms

import (
	"fmt"
	"math"
	"strconv"
)

type Angle struct {
	deg float64
	min float64
	sec float64
}

func NewAngle(deg float64, min float64, sec float64) Angle {
	sign := 1.
	if deg < 0 {
		sign = -1.
	}

	deg, min, sec = math.Abs(deg), math.Abs(min), math.Abs(sec)

	// Normalize floats
	if ideg, fdeg := math.Modf(deg); fdeg != 0 {
		deg = ideg
		min += math.Abs(fdeg) * 60
	}
	if imin, fmin := math.Modf(min); fmin != 0 {
		min = imin
		sec += math.Abs(fmin) * 60
	}

	return Angle{}.Add(Angle{
		deg: deg * sign,
		min: min * sign,
		sec: sec * sign,
	})
}

func NewAngleFromParsed(p Parsed) Angle {
	deg, err := strconv.ParseFloat(p.Deg, 64)
	if err != nil {
		panic(err)
	}
	min, err := strconv.ParseFloat(p.Min, 64)
	if err != nil {
		panic(err)
	}
	sec, err := strconv.ParseFloat(p.Sec, 64)
	if err != nil {
		panic(err)
	}
	switch p.Hemi {
	case NorthType, EastType, "+":
		// good
	case SouthType, WestType, "-":
		deg = deg * -1
	default:
		panic(fmt.Sprintf("invalid hemisphere: %v", p.Hemi))
	}
	return NewAngle(deg, min, sec)
}

func (a Angle) Add(a2 Angle) Angle {
	var carry float64

	sec := a.sec + a2.sec
	carry, a.sec = float64(int(sec/60)), math.Mod(sec, 60)
	min := a.min + a2.min + carry
	carry, a.min = float64(int(min/60)), math.Mod(min, 60)
	a.deg = a.deg + a2.deg + carry
	return a
}

func (a Angle) ToDegrees() float64 {
	return a.deg + (a.min / 60) + (a.sec / 3600)
}

func (a Angle) ToMinutes() float64 {
	return (a.deg * 60) + a.min + (a.sec / 3600)
}

func (a Angle) ToSeconds() float64 {
	return (a.deg * 3600) + (a.min * 60) + a.sec
}

func (a Angle) ToDMS() (deg, min, sec float64) {
	deg, min, sec = a.deg, a.min, a.sec
	return
}
