// Package examples
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package examples

import (
	"fmt"

	"github.com/teocci/go-sorts/src/sortmgr"
)

func Example_strings() {
	groceries := []string{"peppers", "tortillas", "tomatoes", "cheese"}
	sortmgr.Strings(groceries) // or sortutil.Bytes([][]byte)
	fmt.Println(groceries)
	// Output: [cheese peppers tomatoes tortillas]
}