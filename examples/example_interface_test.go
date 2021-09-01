// Package examples
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package examples

import (
	"fmt"
	"github.com/teocci/go-sorts/src/core"
	"github.com/teocci/go-sorts/src/sortmgr"
)

type City struct {
	Name                string
	Latitude, Longitude float32
}

func (c City) String() string { return fmt.Sprintf("%s (%.1f, %.1f)", c.Name, c.Latitude, c.Longitude) }

// ByLatitude implements sort.Interface for []City based on
// the Latitude field, for sorting cities south to north.
type ByLatitude []City

func (a ByLatitude) Len() int      { return len(a) }
func (a ByLatitude) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Float32Key and Float32Less make the sort handle the sign bit and sort NaN
// values to the end.  There are also Float64Key and Float64Less, and
// [Type]Key functions for int types.

// Key returns a uint64 that is lower for more southerly latitudes.
func (a ByLatitude) Key(i int) uint64 {
	return sortmgr.Float32Key(a[i].Latitude)
}
func (a ByLatitude) Less(i, j int) bool {
	return sortmgr.Float32Less(a[i].Latitude, a[j].Latitude)
}

func Example() {
	cities := []City{
		{"Vancouver", 49.3, -123.1},
		{"Tokyo", 35.6, 139.7},
		{"Honolulu", 21.3, -157.8},
		{"Sydney", -33.9, 151.2},
	}

	fmt.Println(cities)
	gosorts.ByUint64(ByLatitude(cities))
	fmt.Println(cities)

	// Output:
	// [Vancouver (49.3, -123.1) Tokyo (35.6, 139.7) Honolulu (21.3, -157.8) Sydney (-33.9, 151.2)]
	// [Sydney (-33.9, 151.2) Honolulu (21.3, -157.8) Tokyo (35.6, 139.7) Vancouver (49.3, -123.1)]
}
