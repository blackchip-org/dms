package dms

type Parsed struct {
	Deg  string
	Min  string
	Sec  string
	Hemi string
}

func NewParsed(deg string, min string, sec string, hemi string) Parsed {
	return Parsed{
		Deg:  deg,
		Min:  min,
		Sec:  sec,
		Hemi: hemi,
	}
}
