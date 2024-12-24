package main

import "math"

// Initializing with uppercase is PUBLIC (Visible in and outside of package)

// Initializing with lowercase is PRIVATE (visible only in package)

// Example:
// Point -> public
// point or _Point -> private

//Pont represents a coordinate in a carthesian plane
type Point struct {
	x float64
	y float64
}

func legs(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return //revision, named return
}

//Distance gets the linear distance between 2 points
func Distance(a, b Point) float64 {
	cx, cy := legs(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}