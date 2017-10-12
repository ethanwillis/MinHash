package main

import (
	"math/rand"
	"time"
)

//-----------------------------------------------------------------------------
func random_dna(n int) string {
	s := make([]byte, n)
	bases := []byte{'A', 'C', 'G', 'T'}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		s[i] = bases[r.Intn(4)]
	}
	return string(s)
}

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
