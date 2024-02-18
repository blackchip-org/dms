package dms_test

import (
	"fmt"

	"github.com/blackchip-org/dms"
)

func Example_example1() {
	p := dms.NewDefaultParser()
	a, err := p.Parse(`1° 3' 6" S`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.6f", a.Degrees())

	// Output:
	// -1.051667
}

func Example_example2() {
	p := dms.NewDefaultParser()
	a, err := p.Parse(`1° 3' 6" S`)
	if err != nil {
		panic(err)
	}
	f := dms.NewFormatter(dms.MinUnit, 3)
	fmt.Println(f.FormatLat(a))

	// Output:
	// 1° 3.100′ S
}
