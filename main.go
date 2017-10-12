package main

import (
	"fmt"
)

//-----------------------------------------------------------------------------
func main() {
	N := 5000
	k, window_size, overlap := 7, 100, 10
	h := NewHash(k, window_size, overlap)
	for i := 0; i < 10; i++ {
		dna := random_dna(N)
		h.Hash(dna, i)
	}
	// h.Print()
	i := 0
	for k, _ := range h.Table {
		fmt.Println(k, h.Lookup(k))
		i++
		if i == 10 {
			break
		}
	}
	key := HashKey{0, 1, 2}
	fmt.Println(h.Lookup(key))
}
