# llrb

GoLLRB is a Left-Leaning Red-Black (LLRB) implementation of 2-3 balanced binary
search trees in Go Language.

## Installation

`go get github.com/wzshiming/llrb`

## Example

This code:

```golang
package main

import (
	"fmt"

	"github.com/wzshiming/llrb"
)

func main() {
	tree := llrb.New()
	for i := 0; i != 10; i++ {
		tree.ReplaceOrInsert(llrb.Int(i))
	}
	tree.DeleteMin()
	tree.Delete(llrb.Int(4))
	fmt.Println(tree)
}
```

prints this output:

```
　　┌>1
　┌>2
　│└>3
─>5
　│┌>6
　└>7
　　│┌>8
　　└>9
```
