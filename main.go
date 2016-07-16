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
// The MIT License (MIT)
//
// Copyright (c) 2016 lexfrei
//
//************************************************************************//

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r int
var scan int

func init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&r, "r", 10000000, "Set the circle size")
	flag.IntVar(&scan, "shots", 10000000, "How many shots will be done")
	flag.Parse()
	fmt.Print("Radius\t = ", r, "\n")
	fmt.Print("Shots\t = ", scan, "\n")
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

//PointInCircle check the point in x,y belonging to the circumference with center in x0, y0 and r=R
func PointInCircle(x, y, x0, y0, R int) bool {
	return (x-x0)*(x-x0)+(y-y0)*(y-y0) <= R*R
}

func main() {
	inS := 0
	start := time.Now()
	for i := 0; i < scan; i++ {
		if PointInCircle(random(-r, r), random(-r, r), 0, 0, r) {
			inS++
		}
	}
	elapsed := time.Since(start)
	pi := float32(inS) * 4 / float32(r) * float32(r) / float32(scan)
	fmt.Printf("Pi\t = %f\n", pi)
	fmt.Printf("Accuracy = %f\n", 1-math.Abs(float64(pi)-math.Pi)/math.Pi)
	fmt.Print("Time\t = ", elapsed, "\n")
}
