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
	return f.format(a, NoAxis)
}

func (f Formatter) FormatLat(a Angle) string {
	return f.format(a, LatAxis)
}

func (f Formatter) FormatLon(a Angle) string {
	return f.format(a, LonAxis)
}

func (f Formatter) format(a Angle, axis Axis) string {
	deg, min, sec := a.DMS()
	sign := 1
	if deg < 0 {
		sign = -1
	}

	var buf strings.Builder
	if axis != NoAxis {
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
	if axis != NoAxis {
		hemi := Hemi(axis, sign)
		fmt.Fprintf(&buf, "%v%v", f.Sep, hemi)
	}
	return buf.String()
}

// func (f Formatter) FormatFields(fields Fields) string {
// 	degSym := firstOf(f.Deg, fields.DegSym)
// 	minSym := firstOf(f.Min, fields.MinSym)
// 	secSym := firstOf(f.Sec, fields.SecSym)

// 	deg := firstOf(fields.Deg, "0")
// 	min := firstOf(fields.Min, "0")
// 	sec := firstOf(fields.Sec, "0")

// 	switch {
// 	case f.To == DegType && f.Places >= 0:
// 		ddeg, err := decimal.NewFromString(deg)
// 		if err != nil {
// 			panic(err)
// 		}
// 		deg = ddeg.StringFixed(int32(f.Places))
// 	case f.To == MinType && f.Places >= 0:
// 		dmin, err := decimal.NewFromString(min)
// 		if err != nil {
// 			panic(err)
// 		}
// 		min = dmin.StringFixed(int32(f.Places))
// 	case f.To == SecType && f.Places >= 0:
// 		dsec, err := decimal.NewFromString(sec)
// 		if err != nil {
// 			panic(err)
// 		}
// 		sec = dsec.StringFixed(int32(f.Places))
// 	}
// 	var buf strings.Builder

// 	if fields.Hemi == "-" {
// 		buf.WriteString(fields.Hemi)
// 	}

// 	buf.WriteString(deg)
// 	buf.WriteString(degSym)

// 	if f.To != DegType {
// 		buf.WriteString(f.Sep)
// 		buf.WriteString(min)
// 		buf.WriteString(minSym)

// 		if f.To != MinType {
// 			buf.WriteString(f.Sep)
// 			buf.WriteString(sec)
// 			buf.WriteString(secSym)
// 		}
// 	}

// 	if fields.Hemi != "" && fields.Hemi != "-" && fields.Hemi != "+" {
// 		buf.WriteString(f.Sep)
// 		buf.WriteString(fields.Hemi)
// 	}
// 	return buf.String()
// }

func firstOf(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}
