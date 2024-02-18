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
	Sep    string
	Places int
	To     Unit
}

func NewFormatter(to Unit, places int) Formatter {
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
	return f.format(a, NoAxis)
}

func (f Formatter) FormatLat(a Angle) string {
	return f.format(a, LatAxis)
}

func (f Formatter) FormatLon(a Angle) string {
	return f.format(a, LonAxis)
}

func (f Formatter) format(a Angle, ax axis) string {
	deg, min, sec := a.DMS()
	sign := 1
	if deg < 0 {
		sign = -1
	}

	var buf strings.Builder
	if ax != NoAxis {
		deg = math.Abs(deg)
	}

	func() {
		if f.To == DegUnit {
			degs := deg + (min / 60) + (sec / 3600)
			if f.Places >= 0 {
				fmt.Fprintf(&buf, "%.*f%v", f.Places, degs, f.Deg)
			} else {
				fmt.Fprintf(&buf, "%v%v", degs, f.Deg)
			}
			return
		}
		fmt.Fprintf(&buf, "%v%v%v", deg, f.Deg, f.Sep)
		if f.To == MinUnit {
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
	if ax != NoAxis {
		hemi := hemi(ax, sign)
		fmt.Fprintf(&buf, "%v%v", f.Sep, hemi)
	}
	return buf.String()
}
