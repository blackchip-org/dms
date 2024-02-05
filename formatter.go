package dms

import (
	"fmt"
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

func (f Formatter) Format(a Angle) string {
	sign, deg, min, sec, err := a.ToFloats()
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	if f.Sign {
		if sign == -1 {
			buf.WriteRune('-')
		}
	}

	func() {
		if f.To == DegType {
			degs := sign * (deg + (min / 60) + (sec / 3600))
			fmt.Fprintf(&buf, "%.*f%v", f.Places, degs, f.Deg)
			return
		}
		fmt.Fprintf(&buf, "%v%v%v", deg, f.Deg, f.Sep)
		if f.To == MinType {
			mins := sign * (min + (sec / 60))
			fmt.Fprintf(&buf, "%.*f%v", f.Places, mins, f.Min)
			return
		}
		fmt.Fprintf(&buf, "%v%v%v%.*f%v", min, f.Min, f.Sep, f.Places, sec, f.Sec)
	}()
	if !f.Sign && a.Hemi != "" {
		fmt.Fprintf(&buf, "%v%v", f.Sep, a.Hemi)
	}
	return buf.String()
}
