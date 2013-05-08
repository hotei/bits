// bits.go (c) 2012 David Rook

/*
This package creates and manipulates arbitrary length bitfields.
  Limitations
  -----------
  Not tested on BIGENDIAN hardware

  Performance
  -----------
  Go 1.1rc2 on 4Ghz AMD/64  using go test -test.bench=".*"
	BenchmarkBitSet-4	500000000     5.69 ns/op
	BenchmarkBitClr-4	500000000     5.89 ns/op
	BenchmarkBitTgl-4	500000000     5.64 ns/op
	BenchmarkBitRead-4	500000000     5.90 ns/op
	BenchmarkBitSetMany-4	100000000    28.50 ns/op  ( sets 4 different bits )
	BenchmarkBitClrMany-4	100000000    29.50 ns/op  ( clrs 4 different bits )
	BenchmarkBitTglMany-4	100000000    28.40 ns/op  ( tgls 4 different bits )

 (c) 2013 David Rook - License is BSD style - see LICENSE.md
  Also see README.md for more info
*/
package bits

import (
	"errors"
	"fmt"
)

//const (
//)

var (
	verbose = false
	license = "bits.go pkg (c) 2013 David Rook released under Simplified BSD License"
)

// BUG(mdr): A few magic numbers related to bits in our choice of uint8.
// My speculation is that uint8 is a better choice than uint64, but that needs
// to be verified through benchmarking a different version.
// This version is very fast as is.
// ? Inspect assembly to see if func call was inlined.

// BUG(mdr): NOTE BENE: This package has NOT been tested on big-endian CPU.
// BUG(mdr): TODO: tests require visual inspection, should compare to expected
//   and call errorf if needed

//  MUSING
// 		its possible log2(8) = 3 may be the best choice vs
// 			bitbox[]uint64 --> log2(64) = 6
//		We frequently do shifts for that many bits as in  byteNum := n >> 3
//			would become wordNum := n >> 6
//		? Does CPU cycles depend on shift repeat factor ?
//		Going to uint64 might break on big-endian since bits are stored
//			in BitField with lsb first.  Speculation for now. Need to test.

var ErrNonEmptySliceRequired = errors.New("bits: Non-empty slice required")

type BitField struct {
	name   string
	maxbit int
	bitbox []uint8
	// falsebits & truebits not quite ready for PrimeTime :-)
	//	falsebits int		use len(bits.FalseBitsLoHi(0,maxbit)) as temp fix?
	//	truebits int		use len(bits.TrueBitsLoHi(0,maxbit)) as temp fix?
}

// func (b *BitField) CountOfFalseBits() int {
// }

// func (b *BitField) CountOfTrueBits() int {
// }

// returns and'ing of all the many bits selected by the slice indices
//   In other words, return false if any bit in index slice is false
func (b *BitField) AndBitsByNdx(many []int) (bool, error) {
	if len(many) <= 0 {
		if verbose {
			fmt.Printf("Warning: AndBitsByNdx(emptySet) isn't meaningful\n")
		}
		return false, ErrNonEmptySliceRequired
	}
	for _, n := range many {
		if b.Bit(n) == false {
			return false, nil
		}
	}
	return true, nil
}

//returns value of the bit at the nth location
func (b *BitField) Bit(n int) bool {
	if n >= b.maxbit {
		for {
			b.bitbox = append(b.bitbox, 0)
			b.maxbit = len(b.bitbox) << 3
			if verbose {
				fmt.Printf("n(%d) required extending maxbit to (%d) \n", n, b.maxbit)
			}
			if b.maxbit > n {
				break
			}
		}
	}
	byteNum := n >> 3
	n -= (byteNum << 3)
	var bitPos uint8 = 1
	bitPos <<= uint(7 - n)
	return !((b.bitbox[byteNum] & bitPos) == 0)
}

// clears the bit at position n
func (b *BitField) ClrBit(n int) {
	if n >= b.maxbit {
		for {
			b.bitbox = append(b.bitbox, 0)
			b.maxbit = len(b.bitbox) << 3
			if verbose {
				fmt.Printf("n(%d) required extending maxbit to (%d) \n", n, b.maxbit)
			}
			if b.maxbit > n {
				break
			}
		}
	}
	byteNum := n >> 3
	n -= (byteNum << 3)
	var bitPos uint8 = 1
	bitPos <<= uint(7 - n)
	//	bitPos = ^bitPos
	//	b.bitbox[byteNum] &= bitPos
	b.bitbox[byteNum] &^= bitPos // &^ is Bit Clear
}

// Clear every bit position selected by the slice of integers
func (b *BitField) ClrBits(many []int) {
	for _, val := range many {
		b.ClrBit(val)
	}
}

// Clear bits in range `[lo..hi]` (inclusive)
func (b *BitField) ClrLoHi(lo, hi int) {
	for i := lo; i <= hi; i++ {
		b.ClrBit(i)
	}
}

// Print bits in range `[lo..hi]` (inclusive) to stdout
// 		not sure about best format - this seems unduly verbose
func (b *BitField) DumpLoHi(lo, hi int) {
	var tmpname string
	if b.name == "" {
		tmpname = "<NoName> "
	} else {
		tmpname = b.name
	}
	fmt.Printf("%s DumpLoHi %d .. %d\n", tmpname, lo, hi)
	for i := lo; i <= hi; i++ {
		tv := b.Bit(i)
		fmt.Printf("\tBit[%d] = %v\n", i, tv)
	}
}

// Print the whole field to stdout
// 		not sure about best format - this seems unduly verbose
func (b *BitField) Dump() {
	b.DumpLoHi(0, b.maxbit)
}

// Returns a slice with indices of the bits with value of false
func (b *BitField) FalseBitsLoHi(lo, hi int) []int {
	rv := make([]int, 0, b.maxbit)
	for i := lo; i <= hi; i++ {
		if !b.Bit(i) {
			rv = append(rv, i)
		}
	}
	return rv
}

// return hex representation of bitfield as string
// 		string will be a multiple of 2 in length
//		see also String() for binary representation
func (b *BitField) HexString() string {
	hs := ""
	for i := 0; i < (b.maxbit / 8); i++ {
		hs += fmt.Sprintf("%02x", b.bitbox[i])
	}
	return hs
}

// returns or'ing of all the bits selected by the slice
//   In other words, return true if any bit in index slice is true
func (b *BitField) OrBitsByNdx(many []int) (bool, error) {
	if len(many) <= 0 {
		if verbose {
			fmt.Printf("Warning: OrBitsByNdx(emptySet) isn't meaningful\n")
		}
		return false, ErrNonEmptySliceRequired
	}
	for _, n := range many {
		if b.Bit(n) == true {
			return true, nil
		}
	}
	return false, nil
}

// Set the bit at position n
func (b *BitField) SetBit(n int) {
	if n >= b.maxbit {
		for {
			b.bitbox = append(b.bitbox, 0)
			b.maxbit = len(b.bitbox) << 3
			if verbose {
				fmt.Printf("n(%d) required extending maxbit to (%d) \n", n, b.maxbit)
			}
			if b.maxbit > n {
				break
			}
		}
	}
	byteNum := n >> 3
	n -= (byteNum << 3)
	var bitPos uint8 = 1
	bitPos <<= uint(7 - n)
	b.bitbox[byteNum] |= bitPos
}

// Set every bit indicated by indices in slice
func (b *BitField) SetBits(many []int) {
	for _, val := range many {
		b.SetBit(val)
	}
}

// Set every bit in the range `[lo..hi]` (inclusive)
func (b *BitField) SetLoHi(lo, hi int) {
	for i := lo; i <= hi; i++ {
		b.SetBit(i)
	}
}

// optional name that will appear on Dump() output
func (b *BitField) SetName(newname string) {
	b.name = newname
}

// just a suggestion, may be overridden during execution
func (b *BitField) SetMaxBitNdx(n int) {
	b.maxbit = n
	b.bitbox = make([]uint8, (n/8)+1)
	//	b.truebits = 0
	//	b.falsebits = maxbit
}

// return binary representation of bitfield as string
// 		string will be a multiple of 8 in length
//		see also HexString() for hex representation
func (b *BitField) String() string {
	rs := ""
	for i := 0; i < (b.maxbit / 8); i++ {
		rs += fmt.Sprintf("%08b", b.bitbox[i])
	}
	return rs
}

// Toggle bit at position n
func (b *BitField) TglBit(n int) {
	if n >= b.maxbit {
		for {
			b.bitbox = append(b.bitbox, 0)
			b.maxbit = len(b.bitbox) << 3
			if verbose {
				fmt.Printf("n(%d) required extending maxbit to (%d) \n", n, b.maxbit)
			}
			if b.maxbit > n {
				break
			}
		}
	}
	byteNum := n >> 3
	n -= (byteNum << 3)
	var bitPos uint8 = 1
	bitPos <<= uint(7 - n)
	b.bitbox[byteNum] ^= bitPos // Xor
}

// Toggle every bit selected by slice indices
func (b *BitField) TglBits(many []int) {
	for _, val := range many {
		b.TglBit(val)
	}
}

// Toggle every bit in the range [lo..hi]
func (b *BitField) TglLoHi(lo, hi int) {
	for i := lo; i <= hi; i++ {
		b.TglBit(i)
	}
}

// return a slice with indices of the bits with value of true
func (b *BitField) TrueBitsLoHi(lo, hi int) []int {
	rv := make([]int, 0, b.maxbit)
	for i := lo; i <= hi; i++ {
		if b.Bit(i) {
			rv = append(rv, i)
		}
	}
	return rv
}
