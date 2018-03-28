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
