// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Bootcamp struct {
// 	Lat  float64
// 	Lon  float64
// 	Date time.Time
// }

// func main() {
// 	fmt.Println(Bootcamp{
// 		Lat:  34.012836,
// 		Lon:  -118.495338,
// 		Date: time.Now(),
// 	})
// }
package main

import (
	"fmt"
	"time"
)

// type Point struct {
// 	X, Y int
// }

// var (
// 	p = Point{1, 2}  // type Point
// 	q = &Point{1, 2} // type *Point
// 	r = Point{X: 1}  // Y:0 is implicit
// 	s = Point{}      //X:0 and Y:0
// )

// func main() {
// 	fmt.Printl
// 	n(p, q, r, s)
// }

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

func main() {
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)
}
