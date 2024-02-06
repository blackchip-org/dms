# dms

A parser and formatter for angles in degrees, minutes, and seconds.

## Parsing

An example of parsing an angle:

```go
	p := dms.NewDefaultParser()
	a, err := p.Parse(`1° 3' 6" S`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.Degrees())

	// Output:
	// -1.051667
```

If a parse result is valid, a `scan.Angle` is returned from the parser that
contains the fields that were extracted. The parser uses the following rules:

- Whitespace is not significant
- A real number by itself is valid: 1, -42, 66.123
- Degrees are not bound to a range. A value of 181 degrees is valid.
- Values in either degrees and minutes (DM) format, or degrees, minutes, and seconds format (DMS), must use unit designators that follow each numeric value
- The unit designator for degrees is either a `°` or `d`
- The unit designator for minutes is either a `'`, `′`, or `m`
- The unit designator for seconds is either a `"`, `″`, or `s`
- A hemisphere designator of `N`, `S`, `E`, `W`, may follow the final unit designator of the value
- Minutes must always be followed degrees, seconds must always be followed by minutes
- Degrees must be an integer when minutes are provided and minutes must be an integer when seconds are provided
- Either a numeric sign (`+` or `-`) or a hemisphere designator may appear, but not both

### Formatting

A formatter is creating with two parameters, the last unit to show (either
`DegType`, `MinType`, or `SecType`), and the number of places for that last
unit. For example, to create a degrees and minutes formatter to 3 digits:

```go
	p := dms.NewDefaultParser()
	a, err := p.Parse(`1° 3' 6" S`)
	if err != nil {
		panic(err)
	}
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.FormatLat(a))

	// Output:
	// 1° 3.100′ S
```

## Status

This package is still a work in progress and is subject to change. If you
find this useful, please drop a line at zc at blackchip dot org.

## License

MIT

