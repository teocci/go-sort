// Package gosorts
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package gosorts

import "sort" // for Interface

// Uint64Interface represents a collection that can be sorted by a uint64
// key.
type Uint64Interface interface {
	sort.Interface
	// Key provides a uint64 key for element i.
	Key(i int) uint64
}

// Int64Interface represents a collection that can be sorted by an int64
// key.
type Int64Interface interface {
	sort.Interface
	// Key provides an int64 key for element i.
	Key(i int) int64
}

// StringInterface represents a collection that can be sorted by a string
// key.
type StringInterface interface {
	sort.Interface
	// Key provides the string key for element i.
	Key(i int) string
}

// BytesInterface represents a collection that can be sorted by a []byte
// key.
type BytesInterface interface {
	sort.Interface
	// Key provides the []byte key for element i.
	Key(i int) []byte
}

// Flip reverses the order of items in a sort.Interface.
func Flip(data sort.Interface) {
	a, b := 0, data.Len()-1
	for b > a {
		data.Swap(a, b)
		a++
		b--
	}
}
