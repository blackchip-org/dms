package dms

import (
	"fmt"
	"strconv"

	"github.com/blackchip-org/scan"
)

type Error struct {
	Pos     scan.Pos
	Message string
}

func NewError(tok scan.Token, format string, args ...any) *Error {
	return &Error{Pos: tok.Pos, Message: fmt.Sprintf(format, args...)}
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.Pos, e.Message)
}

var stateMachine = []func(*scan.Runner, *Fields) (int, error){
	parseSign,       // S0
	parseDegNum,     // S1
	parseDegIntSym,  // S2
	parseRealIntSym, // S3
	parseMinNum,     // S4
	parseSecNum,     // S5
	parseHemi,       // S6
}

type Parser struct {
	ctx     *Context
	scanner scan.Scanner
}

func NewParser(ctx *Context) *Parser {
	return &Parser{ctx: ctx}
}

func NewDefaultParser() *Parser {
	return NewParser(NewContext())
}

func (p *Parser) ParseFields(v string) (Fields, error) {
	var a Fields
	p.scanner.InitFromString("", v)
	r := scan.NewRunner(&p.scanner, p.ctx.RuleSet)

	var state int
	var err error
	for {
		parse := stateMachine[state]
		state, err = parse(r, &a)
		if err != nil {
			return Fields{}, err
		}
		if state == -1 {
			break
		}
	}

	tok := r.This
	if !tok.IsEndOfText() {
		return Fields{}, NewError(tok, "unexpected %v", scan.Quote(tok.Lit))
	}

	return a, nil
}

func (p *Parser) Parse(v string) (Angle, error) {
	parsed, err := p.ParseFields(v)
	if err != nil {
		return Angle{}, err
	}
	deg, err := strconv.ParseFloat(parsed.Deg, 64)
	if err != nil {
		return Angle{}, fmt.Errorf("invalid degrees: %v", parsed.Deg)
	}
	min, err := strconv.ParseFloat(parsed.Min, 64)
	if err != nil {
		return Angle{}, fmt.Errorf("invalid minutes: %v", parsed.Min)
	}
	sec, err := strconv.ParseFloat(parsed.Sec, 64)
	if err != nil {
		return Angle{}, fmt.Errorf("invalid seconds: %v", parsed.Sec)
	}
	switch parsed.Hemi {
	case NorthType, EastType, "+":
		// good
	case SouthType, WestType, "-":
		deg = deg * -1
	default:
		return Angle{}, fmt.Errorf("invalid hemisphere: %v", parsed.Hemi)
	}
	return NewAngle(deg, min, sec), nil
}

// S0
func parseSign(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case "+":
		a.Hemi = "+"
		r.Scan()
	case "-":
		a.Hemi = "-"
		r.Scan()
	}
	return 1, nil
}

// S1
func parseDegNum(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case IntType:
		_, err := strconv.ParseInt(tok.Val, 10, 64)
		if err != nil {
			return -1, NewError(tok, "invalid degree %v", scan.Quote(tok.Lit))
		}
		a.Deg = tok.Val
		r.Scan()
		return 2, nil
	case RealType:
		_, err := strconv.ParseFloat(tok.Val, 64)
		if err != nil {
			return -1, NewError(tok, "invalid degree %v", scan.Quote(tok.Lit))
		}
		a.Deg = tok.Val
		r.Scan()
		return 3, nil
	}
	return -1, NewError(tok, "expected degree, got %v", scan.Quote(tok.Lit))
}

// S2
func parseDegIntSym(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case DegType:
		a.DegSym = tok.Val
		r.Scan()
		return 4, nil
	}
	return -1, nil
}

// S3
func parseRealIntSym(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case DegType:
		a.DegSym = tok.Val
		r.Scan()
		return 6, nil
	}
	return -1, nil
}

// S4
func parseMinNum(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case IntType:
		min, err := strconv.ParseInt(tok.Val, 10, 64)
		if err != nil || min >= 60 || min <= -60 {
			return -1, NewError(tok, "invalid minute %v", scan.Quote(tok.Lit))
		}
		a.Min = tok.Val

		tok := r.Scan()
		if tok.Type != MinType {
			return -1, NewError(tok, "expected minute symbol, got %v", scan.Quote(tok.Lit))
		}
		a.MinSym = tok.Val
		r.Scan()
		return 5, nil
	case RealType:
		min, err := strconv.ParseFloat(tok.Val, 64)
		if err != nil || min >= 60 || min <= -60 {
			return -1, NewError(tok, "invalid minute %v", scan.Quote(tok.Lit))
		}
		a.Min = tok.Val

		tok := r.Scan()
		if tok.Type != MinType {
			return -1, NewError(tok, "expected minute symbol, got %v", scan.Quote(tok.Lit))
		}
		a.MinSym = tok.Val
		r.Scan()
		return 6, nil
	}
	return 6, nil
}

// S5
func parseSecNum(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	switch tok.Type {
	case IntType:
		sec, err := strconv.ParseInt(tok.Val, 10, 64)
		if err != nil || sec >= 60 || sec <= -60 {
			return -1, NewError(tok, "invalid second %v", scan.Quote(tok.Lit))
		}
		a.Sec = tok.Val
		r.Scan()
	case RealType:
		sec, err := strconv.ParseFloat(tok.Val, 64)
		if err != nil || sec >= 60 || sec <= -60 {
			return -1, NewError(tok, "invalid second %v", scan.Quote(tok.Lit))
		}
		a.Sec = tok.Val
		r.Scan()
	default:
		return 6, nil
	}

	tok = r.This
	if tok.Type != SecType {
		return -1, NewError(tok, "expected second symbol, got %v", scan.Quote(tok.Lit))
	}
	a.SecSym = tok.Val
	r.Scan()

	return 6, nil
}

// S6
func parseHemi(r *scan.Runner, a *Fields) (int, error) {
	tok := r.This
	var hemi string
	switch tok.Type {
	case NorthType:
		hemi = NorthType
		r.Scan()
	case SouthType:
		hemi = SouthType
		r.Scan()
	case EastType:
		hemi = EastType
		r.Scan()
	case WestType:
		hemi = WestType
		r.Scan()
	}

	if hemi != "" && a.Hemi != "" {
		return -1, NewError(tok, "only one of %v or %v are allowed", scan.Quote(a.Hemi), scan.Quote(hemi))
	}
	if hemi != "" {
		a.Hemi = hemi
	}
	return -1, nil
}
