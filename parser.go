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

var stateMachine = []func(*scan.Runner, *Angle) (int, error){
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

func (p *Parser) Parse(v string) (Angle, error) {
	var a Angle
	p.scanner.InitFromString("", v)
	r := scan.NewRunner(&p.scanner, p.ctx.RuleSet)
	first := r.This

	var state int
	var err error
	for {
		parse := stateMachine[state]
		state, err = parse(r, &a)
		if err != nil {
			return Angle{}, err
		}
		if state == -1 {
			break
		}
	}

	if a.Sign != "" && a.Hemi != "" {
		return Angle{}, NewError(first, `only one of %v or %v allowed`, scan.Quote(a.Sign), scan.Quote(a.Hemi))
	}

	tok := r.This
	if !tok.IsEndOfText() {
		return Angle{}, NewError(tok, "unexpected %v", scan.Quote(tok.Lit))
	}

	return a, nil
}

// S0
func parseSign(r *scan.Runner, a *Angle) (int, error) {
	tok := r.This
	switch tok.Type {
	case "+":
		a.Sign = "+"
		r.Scan()
	case "-":
		a.Sign = "-"
		r.Scan()
	}
	return 1, nil
}

// S1
func parseDegNum(r *scan.Runner, a *Angle) (int, error) {
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
func parseDegIntSym(r *scan.Runner, a *Angle) (int, error) {
	tok := r.This
	switch tok.Type {
	case DegType:
		r.Scan()
		return 4, nil
	}
	return -1, nil
}

// S3
func parseRealIntSym(r *scan.Runner, a *Angle) (int, error) {
	tok := r.This
	switch tok.Type {
	case DegType:
		r.Scan()
		return 6, nil
	}
	return -1, nil
}

// S4
func parseMinNum(r *scan.Runner, a *Angle) (int, error) {
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
		r.Scan()
		return 6, nil
	}
	return 6, nil
}

// S5
func parseSecNum(r *scan.Runner, a *Angle) (int, error) {
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
	r.Scan()

	return 6, nil
}

// S6
func parseHemi(r *scan.Runner, a *Angle) (int, error) {
	tok := r.This
	switch tok.Type {
	case NorthType:
		a.Hemi = NorthType
		r.Scan()
	case SouthType:
		a.Hemi = SouthType
		r.Scan()
	case EastType:
		a.Hemi = EastType
		r.Scan()
	case WestType:
		a.Hemi = WestType
		r.Scan()
	}
	return -1, nil
}
