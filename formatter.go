package dms

import (
	"fmt"
	"math"
	"strings"
)

type axis int

const (
	noAxis axis = iota
	latAxis
	lonAxis
)

type Formatter struct {
	Deg    string
	Min    string
	Sec    string
	Sign   bool
	Sep    string
	Places int
	To     string
}

func NewFormatter(to string, places int) Formatter {
	return Formatter{
		Deg:    "°",
		Min:    "′",
		Sec:    "″",
		Sep:    " ",
		Places: places,
		To:     to,
	}
}

func (f Formatter) WithSymbols(deg string, min string, sec string) Formatter {
	f.Deg, f.Min, f.Sec = deg, min, sec
	return f
}

func (f Formatter) WithSep(sep string) Formatter {
	f.Sep = sep
	return f
}

func (f Formatter) Format(a Angle) string {
	return f.format(a, noAxis)
}

func (f Formatter) FormatLat(a Angle) string {
	return f.format(a, latAxis)
}

func (f Formatter) FormatLon(a Angle) string {
	return f.format(a, lonAxis)
}

func (f Formatter) format(a Angle, axis axis) string {
	deg, min, sec := a.DMS()
	sign := 1
	if deg < 0 {
		sign = -1
	}

	var buf strings.Builder
	if axis != noAxis {
		deg = math.Abs(deg)
	}

	func() {
		if f.To == DegType {
			degs := deg + (min / 60) + (sec / 3600)
			if f.Places >= 0 {
				fmt.Fprintf(&buf, "%.*f%v", f.Places, degs, f.Deg)
			} else {
				fmt.Fprintf(&buf, "%v%v", degs, f.Deg)
			}
			return
		}
		fmt.Fprintf(&buf, "%v%v%v", deg, f.Deg, f.Sep)
		if f.To == MinType {
			mins := math.Abs(min + (sec / 60))
			if f.Places >= 0 {
				fmt.Fprintf(&buf, "%.*f%v", f.Places, mins, f.Min)
			} else {
				fmt.Fprintf(&buf, "%v%v", mins, f.Min)
			}
			return
		}
		mins := math.Abs(min)
		secs := math.Abs(sec)
		if f.Places >= 0 {
			fmt.Fprintf(&buf, "%v%v%v%.*f%v", mins, f.Min, f.Sep, f.Places, secs, f.Sec)
		} else {
			fmt.Fprintf(&buf, "%v%v%v%v%v", mins, f.Min, f.Sep, secs, f.Sec)
		}
	}()
	if axis != noAxis {
		var hemi string
		switch {
		case sign >= 0 && axis == latAxis:
			hemi = NorthType
		case sign < 0 && axis == latAxis:
			hemi = SouthType
		case sign >= 0 && axis == lonAxis:
			hemi = EastType
		case sign < 0 && axis == lonAxis:
			hemi = WestType
		default:
			panic("unreachable")
		}
		fmt.Fprintf(&buf, "%v%v", f.Sep, hemi)
	}
	return buf.String()

}
