# dms 

A parser and formatter for angles in degrees, minutes, and seconds.

## Parsing

An example of parsing an angle:

```go 
	p := NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
```

