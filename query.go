//-----------------------------------------------------------------------------
// Min hash DNA sequences
// Author: Vinhthuy Phan, 2017
//-----------------------------------------------------------------------------
package main

import (
	"fmt"
	"sort"
)

//-----------------------------------------------------------------------------
// Find the genome(s) the read or its reverse complement belongs to.
//-----------------------------------------------------------------------------
func (h *Hash) Query(read string) {
	kmer := read[0:h.K]
	last := h.K - 1
	largest_exp := 1 << (2 * uint(h.K-1))

	// value : numerical value of the current kmer
	values := make([]int, 0)
	value := kmer_to_dec(kmer)
	values = append(values, value)

	// value_rc : numerical value of the reverse complement of the current kmer
	values_rc := make([]int, 0)
	value_rc := kmer_to_dec_rc(kmer)
	values_rc = append(values_rc, value_rc)

	for i := 1; i < len(read)-h.K+1; i++ {
		// fmt.Println(
		// 	value,
		// 	value_rc,
		// 	read[i-1:i+h.K-1],
		// 	kmer_to_dec(read[i-1:i+h.K-1]),
		// 	kmer_to_dec_rc(read[i-1:i+h.K-1]),
		// )
		last += 1
		value = (value-base_to_num(read[i-1])*largest_exp)*4 + base_to_num(read[last])
		values = append(values, value)
		value_rc = (value_rc-base_to_num_rc(read[i-1]))/4 + base_to_num_rc(read[last])*largest_exp
		values_rc = append(values_rc, value_rc)
	}
	// fmt.Println(value, value_rc)
	sort.Ints(values)
	sort.Ints(values_rc)
	fmt.Println("MinHash values of      ", read, "is", values[0:M])
	fmt.Println("MinHash values of rc of", read, "is", values_rc[0:M])
}
