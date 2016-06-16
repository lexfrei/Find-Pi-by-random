//************************************************************************//
// Find Pi by Random
//
// The number π is a mathematical constant, the ratio of a circle's
// circumference to its diameter, commonly approximated as 3.14159
//
// This program makes the "shots" in the square and finds hits in relation
// to all the inscribed circle "shots"
//
// Usage:
// program [options]
//
// Options:
// -r=INT
//   set circle size
// -shots=INT
//   set shots count
//
//************************************************************************//

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var r int
var scan int

func init() {
	flag.IntVar(&r, "r", 10000, "Set the circle size")
	flag.IntVar(&scan, "shots", 10000, "How many shots will be done")
	flag.Parse()
	fmt.Print("Radius\t= ", r, "\n")
	fmt.Print("Shots\t= ", scan, "\n")
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

//PointInCircle проверяет, есть ли жизнь на марсе
func PointInCircle(x, y, x0, y0, R int) bool {
	return (x-x0)*(x-x0)+(y-y0)*(y-y0) <= R*R
}

func main() {
	rand.Seed(time.Now().Unix())
	inS := 0
	for i := 0; i < scan; i++ {
		if PointInCircle(random(-r, r), random(-r, r), 0, 0, r) {
			inS++
		}
	}
	pi := float32(inS) * 4 / float32(r) * float32(r) / float32(scan)
	fmt.Printf("Pi\t= %f\n", float32(pi))
}
