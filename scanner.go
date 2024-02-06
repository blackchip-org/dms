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
	SignRule  = scan.NewClassRule(scan.IsSign)
	DegRule   = scan.NewClassRule(scan.Rune('d', '°')).WithType(DegType)
	MinRule   = scan.NewClassRule(scan.Rune('m', '\'', '′')).WithType(MinType)
	SecRule   = scan.NewClassRule(scan.Rune('s', '"', '″')).WithType(SecType)
	EastRule  = scan.NewClassRule(scan.Rune('E')).WithType(EastType)
	NorthRule = scan.NewClassRule(scan.Rune('N')).WithType(NorthType)
	SouthRule = scan.NewClassRule(scan.Rune('S')).WithType(SouthType)
	WestRule  = scan.NewClassRule(scan.Rune('W')).WithType(WestType)
)

type Context struct {
	RuleSet scan.RuleSet
}

func NewContext() *Context {
	c := &Context{}
	c.RuleSet = scan.NewRuleSet(
		scan.SkipSpaceRule,
		scan.RealRule,
		SignRule,
		DegRule, MinRule, SecRule,
		EastRule, NorthRule, SouthRule, WestRule,
	)
	return c
}
