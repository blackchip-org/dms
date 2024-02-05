package dms_test

import (
	"fmt"

	"github.com/blackchip-org/dms"
)

func Example_example1() {
	p := dms.NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
}

func Example_example2() {
	a := dms.NewAngle("", "1", "3", "6", "S")
	fmt.Printf("%.6f", a.ToDegrees())

	// Output:
	// -1.051667
}

func Example_example3() {
	p := dms.NewDefaultParser()
	a, err := p.Parse("1° 3′ 6″ S")
	if err != nil {
		panic(err)
	}
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.Format(a))

	// Output:
	// 1° 3.100′ S
}

func Example_example4() {
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
}

func Example_example5() {
	f := dms.NewFormatter(dms.MinType, 3)
	fmt.Println(f.FormatLat(-1.5, 0, 0))
	fmt.Println(f.FormatLon(-1, 30, 30))

	// Output:
	// 1° 30.000′ S
	// 1° 30.500′ W
}
