//-----------------------------------------------------------------------------
// Min hash DNA sequences
// Author: Vinhthuy Phan, 2017
//-----------------------------------------------------------------------------
package main

import (
	"fmt"
	"sort"
)

// M is the number of hash values
const M = 3

// hash_key is the M smallest values in the window.
type HashKey [M]int

// hash_value is (sequence id, window (starting position))
type HashVal [2]int

type Hash struct {
	K       int
	Window  int
	Overlap int
	Table   map[HashKey][]HashVal
}

//-----------------------------------------------------------------------------
func NewHash(k, window, overlap int) *Hash {
	h := &Hash{K: k, Window: window, Overlap: overlap}
	h.Table = make(map[HashKey][]HashVal)
	return h
}

//-----------------------------------------------------------------------------
func (h *Hash) Lookup(key HashKey) []HashVal {
	return h.Table[key]
}

//-----------------------------------------------------------------------------
// M smallest "k"-mers in the "window" start at "starting_pos"
//-----------------------------------------------------------------------------
func (h *Hash) HashWindow(dna string, id, starting_pos int) {
	if starting_pos+h.K-1 > len(dna)-1 {
		return
	}
	kmer := dna[starting_pos : starting_pos+h.K]
	last := starting_pos + h.K - 1
	largest_exp := 1 << (2 * uint(h.K-1))
	values := make([]int, 0)
	value := kmer_to_dec(kmer)
	values = append(values, value)
	for i := starting_pos + 1; i <= starting_pos+h.Window && i < len(dna)-h.K+1; i++ {
		last += 1
		value -= base_to_num(dna[i-1]) * largest_exp
		value = value*4 + base_to_num(dna[last])
		values = append(values, value)
	}
	sort.Ints(values)
	// hash_key is the M smallest values in the window.
	// hash_value is (sequence id, window (starting position))
	hash_key := HashKey{}
	hash_value := HashVal{id, starting_pos}
	for i := 0; i < M; i++ {
		hash_key[i] = values[i]
	}
	h.Table[hash_key] = append(h.Table[hash_key], hash_value)
}

//-----------------------------------------------------------------------------
// compute all k-mer hashes in a window of size w with given overlap.
//-----------------------------------------------------------------------------
func (h *Hash) Hash(seq string, seq_id int) {
	n := len(seq)
	start := 0
	end := n - h.K + 1
	for p := start; p <= end; p = p + h.Window - h.Overlap {
		h.HashWindow(seq, seq_id, p)
	}
}

//-----------------------------------------------------------------------------
func (h *Hash) Print() {
	fmt.Printf("K: %d, Window: %d, M: %d, Overlap: %d\n",
		h.K, h.Window, M, h.Overlap)
	for key, value := range h.Table {
		fmt.Println("Key:", key, "Hash values:", value)
	}
}

//-----------------------------------------------------------------------------
