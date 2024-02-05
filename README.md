# dms

A parser and formatter for angles in degrees, minutes, and seconds.

## Parsing

An example of parsing an angle:

```go
	p := dms.NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.ToDegrees())

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

If fields have already been extracted, the `NewAngle()` function can be used
instead:

```go
	a := dms.NewAngle("", "1", "3", "6", "S")
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
```

### Formatting

A formater is creating with two parameters, the last unit to show (either
`DegType`, `MinType`, or `SecType`), and the number of places for that last
unit. For example, to create a degrees and minutes formatter to 3 digits:

```go
	p := dms.NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.Format(a))

	// Output:
	// 1° 3.100′ S
```

For signed values, the `AsLat()` and `AsLon()` methods can be used to provide
additional context:

```go
	p := dms.NewDefaultParser()
	a, err := p.Parse("-1° 3′ 6″")
	if err != nil {
		panic(err)
	}
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.Format(a.AsLat()))
	fmt.Println(f.Format(a.AsLon()))

	// Output:
	// 1° 3.100′ S
	// 1° 3.100′ W
```

Float values can formatted with `FormatLat` and `FormatLon`. The degrees
value should contain the sign of the overall value. Signs are ignored on
minutes and degree values:

```go
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.FormatLat(-1.5, 0, 0))
	fmt.Println(f.FormatLon(-1, 30, 30))

	// Output:
	// 1° 30.000′ S
	// 1° 30.500′ W
```

## Status

This package is still a work in progress and is subject to change. If you
find this useful, please drop a line at zc at blackchip dot org.

## License

MIT

