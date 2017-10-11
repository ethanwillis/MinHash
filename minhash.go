package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//-----------------------------------------------------------------------------
func base_to_num(base byte) int {
	if base == 'A' {
		return 0
	} else if base == 'C' {
		return 1
	} else if base == 'G' {
		return 2
	} else if base == 'T' {
		return 3
	} else {
		panic("Unknown character: " + string(base))
	}
}

//-----------------------------------------------------------------------------
func kmer_to_dec(kmer string) int {
	value := 0
	exp := 1
	for i := len(kmer) - 1; i >= 0; i-- {
		if kmer[i] == 'A' {
			value += 0 * exp
		} else if kmer[i] == 'C' {
			value += 1 * exp
		} else if kmer[i] == 'G' {
			value += 2 * exp
		} else if kmer[i] == 'T' {
			value += 3 * exp
		} else {
			panic("Unknown character: " + string(kmer[i]))
		}
		exp *= 4
	}
	return value
}

//-----------------------------------------------------------------------------
// return "how_many" smallest "k"-mers in the "window" start at "starting_pos"
//-----------------------------------------------------------------------------
func min_hash(dna string, starting_pos, k, window, how_many int) []int {
	if starting_pos+k-1 > len(dna)-1 {
		return nil
	}
	kmer := dna[starting_pos : starting_pos+k]
	last := starting_pos + k - 1
	largest_exp := 1 << (2 * uint(k-1))
	values := make([]int, 0)
	value := kmer_to_dec(kmer)
	values = append(values, value)
	// fmt.Println(kmer, value, largest_exp)
	for i := starting_pos + 1; i <= starting_pos+window && i < len(dna)-k+1; i++ {
		last += 1
		value -= base_to_num(dna[i-1]) * largest_exp
		value = value*4 + base_to_num(dna[last])
		// fmt.Println(i, dna[i:last+1], dna[i:i+k], value, kmer_to_dec(dna[i:i+k]))
		values = append(values, value)
	}
	sort.Ints(values)
	// fmt.Println(values)
	return values[0:how_many]
}

//-----------------------------------------------------------------------------
// compute all k-mer hashes in a window of size w with given overlap.
//-----------------------------------------------------------------------------
func hash_dna(seq string, k, w, overlap int) {
	n := len(seq)
	start := 0
	end := n - k + 1
	how_many := 3
	for p := start; p <= end; p = p + w - overlap {
		h := min_hash(seq, p, k, w, how_many)
		fmt.Println("start,end,min_hash: ", p, p+w, h)
	}
}

//-----------------------------------------------------------------------------
func random_dna(n int) string {
	s := make([]byte, n)
	bases := []byte{'A', 'C', 'G', 'T'}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		base := bases[r.Intn(4)]
		s[i] = base
	}
	return string(s)
}

//-----------------------------------------------------------------------------
func main() {
	N := 500
	k := 7
	w := 100
	overlap := 10
	fmt.Println(N, k, w, overlap)
	dna := random_dna(N)
	hash_dna(dna, k, w, overlap)
}
