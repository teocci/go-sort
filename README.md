## go-sort [![Go Reference][1]][2]
`go-sort` is an open-source tool for sorting. It provides parallel radix sorting by a `string`, `[]byte`, or `(u)int64` key, and a parallel `Quicksort(data)`. `sortmgr` sorts common slice types and adds functions to help sort `floats`.


## Features
Usually, stick to `stdlib sort`: that's fast, standard, and simpler. But this package may help if sorting huge datasets is a bottleneck for you. To get a sense of the potential gains, some timings are available.

To radix sort, implement `sort.Interface` plus one more method, `Key(i int)`, returning the key for an item as `string`/`[]byte`/`(u)int64`, and call `sorts.ByString`, `ByBytes`, `ByUint64`, or `ByInt64`.

There's no `Reverse()`, but `sorts.Flip(data)` will flip ascending-sorted data to descending. There's no stable sort. The string sorts just compare byte values; `é` won't sort next to `e`. Set sorts.MaxProcs if you want to limit concurrency. The package checks that data is sorted after every run and panics(!) if not.

Credit (but no blame, or claim of endorsement) to the authors of stdlib sort; this uses its `qSort`, `tests`, and `interface`, and the clarity of the code helped make this possible.

## Usage
```go
package main

import (
	"fmt"
	"github.com/teocci/go-sorts/src/core"
	"github.com/teocci/go-sorts/src/sortmgr"
)

func main() {
	scores := []int{39, 492, 4912, 39, -10, 4, 92}
	
	data := sortmgr.IntSlice(scores)
	data.Sort()
	
	gosorts.Flip(data) // high scores first
	
	fmt.Println(scores)
	// Output: [4912 492 92 39 39 4 -10]
}
```

----
#### In case you face trouble, please feel free to open an issue.

[1]: https://pkg.go.dev/badge/github.com/teocci/go-sort.svg
[2]: https://pkg.go.dev/github.com/teocci/go-sort
[3]: https://github.com/teocci/go-sort/releases/latest