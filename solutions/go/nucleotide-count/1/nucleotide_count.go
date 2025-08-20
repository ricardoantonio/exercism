package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

func newHistogram() Histogram {
	return Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
}

// addNucleotide increments the count of a nucleotide in the histogram.
// returns an error if n is not a valid nucleotide ('A', 'C', 'G', 'T').
func (h Histogram) addNucleotide(n rune) error {
	switch n {
	case 'A', 'C', 'G', 'T':
		h[n]++
		return nil
	default:
		return errors.New("invalid nucleotide: " + string(n))
	}
}

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := newHistogram()
	for _, n := range d {
		err := h.addNucleotide(n)
		if err != nil {
			return nil, err
		}

	}
	return h, nil
}
