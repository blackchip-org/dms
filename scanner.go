package dms

import "github.com/blackchip-org/scan"

const (
	IntType   = scan.IntType
	RealType  = scan.RealType
	DegType   = "deg"
	MinType   = "min"
	SecType   = "sec"
	EastType  = "E"
	NorthType = "N"
	SouthType = "S"
	WestType  = "W"
)

var (
	Sign  = scan.NewClassRule(scan.Sign)
	Deg   = scan.NewClassRule(scan.Rune('d', '°')).WithType(DegType)
	Min   = scan.NewClassRule(scan.Rune('m', '\'', '′')).WithType(MinType)
	Sec   = scan.NewClassRule(scan.Rune('s', '"', '″')).WithType(SecType)
	East  = scan.NewClassRule(scan.Rune('E')).WithType(EastType)
	North = scan.NewClassRule(scan.Rune('N')).WithType(NorthType)
	South = scan.NewClassRule(scan.Rune('S')).WithType(SouthType)
	West  = scan.NewClassRule(scan.Rune('W')).WithType(WestType)
)

type Context struct {
	RuleSet scan.RuleSet
}

func NewContext() *Context {
	c := &Context{}
	c.RuleSet = scan.NewRuleSet(
		scan.NewSpaceRule(scan.Whitespace),
		scan.Real,
		Sign,
		Deg, Min, Sec,
		East, North, South, West,
	)
	return c
}
