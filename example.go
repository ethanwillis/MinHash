//-----------------------------------------------------------------------------
// Min hash DNA sequences
// Author: Vinhthuy Phan, 2017
//-----------------------------------------------------------------------------
package main

import (
	"fmt"
)

//-----------------------------------------------------------------------------
func main() {
	N := 5000
	k, window_size, overlap := 5, 100, 10
	h := NewHash(k, window_size, overlap)
	for i := 0; i < 10; i++ {
		fmt.Println("Hash random dna of length", N, "with id", i)
		dna := random_dna(N)
		h.Hash(dna, i)
	}
	// h.Print()
	for i := 0; i < 20; i++ {
		read := random_dna(50)
		h.Query(read)
	}
	// fmt.Println("Look up a few keys.")
	// i := 0
	// for k, _ := range h.Table {
	// 	fmt.Println("Key:", k, "Value:", h.Lookup(k))
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// }
	// key := HashKey{0, 1, 2}
	// fmt.Println("Key:", key, "Value:", h.Lookup(key))
}
