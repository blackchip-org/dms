package dms

import (
	"fmt"
	"math"
	"strings"
)

type Formatter struct {
	Deg    string
	Min    string
	Sec    string
	Sign   bool
	North  string
	South  string
	West   string
	East   string
	Sep    string
	Places int
	To     string
}

func NewFormatter(to string, places int) Formatter {
	return Formatter{
		Deg:    "°",
		Min:    "′",
		Sec:    "″",
		Sign:   false,
		Sep:    " ",
		Places: places,
		To:     to,
	}
}

func (f Formatter) WithSymbols(deg string, min string, sec string) Formatter {
	f.Deg, f.Min, f.Sec = deg, min, sec
	return f
}

func (f Formatter) WithSign(s bool) Formatter {
	f.Sign = s
	return f
}

func (f Formatter) WithSep(sep string) Formatter {
	f.Sep = sep
	return f
}

func (f Formatter) format(sign float64, deg float64, min float64, sec float64, hemi string) string {
	var buf strings.Builder
	if f.Sign {
		if sign == -1 || hemi == SouthType || hemi == WestType {
			buf.WriteRune('-')
		}
	}

	deg, min, sec = normalizeFloats(deg, min, sec)
	func() {
		if f.To == DegType {
			degs := deg + (min / 60) + (sec / 3600)
			fmt.Fprintf(&buf, "%.*f%v", f.Places, degs, f.Deg)
			return
		}
		fmt.Fprintf(&buf, "%v%v%v", deg, f.Deg, f.Sep)
		if f.To == MinType {
			mins := min + (sec / 60)
			fmt.Fprintf(&buf, "%.*f%v", f.Places, mins, f.Min)
			return
		}
		fmt.Fprintf(&buf, "%v%v%v%.*f%v", min, f.Min, f.Sep, f.Places, sec, f.Sec)
	}()
	if !f.Sign && hemi != "" {
		fmt.Fprintf(&buf, "%v%v", f.Sep, hemi)
	}
	return buf.String()
}

func (f Formatter) Format(a Angle) string {
	sign, deg, min, sec, err := a.ToFloats()
	if err != nil {
		panic(err)
	}
	return f.format(sign, deg, min, sec, a.Hemi)
}

func (f Formatter) FormatLat(d float64, m float64, s float64) string {
	hemi := NorthType
	sign := float64(1)
	if d < 0 {
		hemi = SouthType
		sign = -1
	}
	return f.format(sign, math.Abs(d), math.Abs(m), math.Abs(s), hemi)
}

func (f Formatter) FormatLon(d float64, m float64, s float64) string {
	hemi := EastType
	sign := float64(1)
	if d < 0 {
		hemi = WestType
		sign = -1
	}
	return f.format(sign, math.Abs(d), math.Abs(m), math.Abs(s), hemi)
}
