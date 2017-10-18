package main

import (
  "bufio"
  "bytes"
  "io"
)

type SeqRecord struct {
	Name, Seq, Qual string
}

type FastqReader struct {
	Reader          *bufio.Reader
	last, seq, qual []byte 
	finished        bool
	rec             record
}

func (fq *FastqReader) iterLines() ([]byte, bool) {
	line, err := fq.Reader.ReadSlice('\n')
  	if err != nil {
    	if err == io.EOF {
      		return line, true
    	} else {
      		panic(err)
    	}
  	}
  	return line, false
}


func (fq *FastqReader) Iter() (record, bool) {
	if fq.finished {
		return fq.rec, fq.finished
	}
  	
  	// Read the seq id (fasta or fastq)
  	if fq.last == nil {
    	for l, done := fq.iterLines(); !done; l, done = fq.iterLines() {
      		if l[0] == '>' || l[0] == '@' { // read id
        		fq.last = l[0 : len(l)-1]
        		break
      		}
    	}

    	if fq.last == nil { // We couldn't find a valid record, no more data in file
      		fq.finished = true
      		return fq.rec, fq.finished
    	}
  	}	
  
  	fq.rec.Name = string(bytes.SplitN(fq.last, []byte(" "), 1)[0])
  	fq.last = nil

  	// Now read the sequence
  	fq.seq = fq.seq[:0]
 	for l, done := fq.iterLines(); !done; l, done = fq.iterLines() {
    	c := l[0]
    	if c == '+' || c == '>' || c == '@' {
      		fq.last = l[0 : len(l)-1]
      		break
    	}
    	fq.seq = append(fq.seq, l[0:len(l)-1]...)
  	}
  	fq.rec.Seq = string(fq.seq)

  	if fq.last != nil { 
    	if fq.last[0] != '+' { 
      		return fq.rec, fq.finished
    	}
    
    	leng := 0
    	fq.qual = fq.qual[:0]
    
    	for l, done := fq.iterLines(); !done; l, done = fq.iterLines() {
    		fq.qual = append(fq.qual, l[0:len(l)-1]...)
      		leng += len(l)
      		if leng >= len(fq.seq) { 
        		fq.last = nil
        		fq.rec.Qual = string(fq.qual)
        		return fq.rec, fq.finished
      		}
    	}
    
    	fq.finished = true
    	fq.rec.Qual = string(fq.qual)
  	}
  	return fq.rec, fq.finished 
}