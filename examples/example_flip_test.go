// Package examples
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package examples

import (
	"fmt"
	"github.com/teocci/go-sorts/src/core"
	"github.com/teocci/go-sorts/src/sortmgr"
)

func Example_flip() {
	scores := []int{39, 492, 4912, 39, -10, 4, 92}
	data := sortmgr.IntSlice(scores)
	data.Sort()
	gosorts.Flip(data) // high scores first
	fmt.Println(scores)
	// Output: [4912 492 92 39 39 4 -10]
}
