package dms

import (
	"fmt"
	"math"
	"strings"
)

const pi180 = math.Pi / 180.0

type Fields struct {
	Deg    string
	Min    string
	Sec    string
	DegSym string
	MinSym string
	SecSym string
	Hemi   string
}

func (f Fields) Sign() int {
	switch f.Hemi {
	case SouthType, WestType, "-":
		return -1
	}
	return 1
}

func (f Fields) IsDD() bool {
	return f.Min == "" && f.Sec == ""
}

func (f Fields) IsDM() bool {
	return f.Min != "" && f.Sec == ""
}

func (f Fields) IsDMS() bool {
	return f.Min != "" && f.Sec != ""
}

func (f Fields) String() string {
	var buf strings.Builder

	sign := ""
	if f.Hemi == "-" || f.Hemi == "+" {
		sign = f.Hemi
	}
	degSym := f.DegSym
	if degSym == "" {
		degSym = "°"
	}
	fmt.Fprintf(&buf, "%v%v%v", sign, f.Deg, degSym)

	if f.Min != "" {
		minSym := f.MinSym
		if minSym == "" {
			minSym = "′"
		}
		fmt.Fprintf(&buf, " %v%v", f.Min, minSym)
	}

	if f.Sec != "" {
		secSym := f.SecSym
		if secSym == "" {
			secSym = "″"
		}
		fmt.Fprintf(&buf, " %v%v", f.Sec, secSym)
	}

	if sign == "" && f.Hemi != "" {
		fmt.Fprintf(&buf, " %v", f.Hemi)
	}

	return buf.String()
}

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

func (a Angle) String() string {
	return fmt.Sprintf("(%v,%v,%v)", a.deg, a.min, a.sec)
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

func (a Angle) Sub(a2 Angle) Angle {
	a2neg := Angle{deg: -a2.deg, min: -a2.min, sec: -a2.sec}
	return a.Add(a2neg)
}

func (a Angle) Degrees() float64 {
	return a.deg + (a.min / 60) + (a.sec / 3600)
}

func (a Angle) Minutes() float64 {
	return (a.deg * 60) + a.min + (a.sec / 60)
}

func (a Angle) Seconds() float64 {
	return (a.deg * 3600) + (a.min * 60) + a.sec
}

func (a Angle) Radians() float64 {
	return a.Degrees() * pi180
}

func (a Angle) DMS() (deg, min, sec float64) {
	deg, min, sec = a.deg, math.Abs(a.min), math.Abs(a.sec)
	return
}
