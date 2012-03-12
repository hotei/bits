// bits.go

// bits go pkg (c) 2012 David Rook
// please look for text string "BUG" as there are a few potential lurkers
package bits

import (
	"fmt"
	"errors"
)

const (
)
var (
	verbose = false
	license = "bits go pkg (c) 2012 David Rook released under Simplified BSD License"
)

type BitField struct {
	name   string
	maxbit int
	bitbox []uint8
	// not quite ready for PrimeTime - need to bench it first
//	falsebits int		use len(bits.FalseBitsLoHi(0,maxbit)) as temp fix?
//	truebits int		use len(bits.TrueBitsLoHi(0,maxbit)) as temp fix?
}

var errNonEmptySliceRequired = errors.New("bits: Non-empty slice required")
// BUG(mdr): AndBitsByNdx() what to report if fed empty set? false now

// returns and'ing of all the many bits selected by the slice indices
func (b *BitField) AndBitsByNdx(many []int) (bool, error) {
	if len(many) <= 0 {
		if verbose { fmt.Printf("Warning: AndBitsByNdx(emptySet) isn't meaningful\n")}
		return false, errNonEmptySliceRequired 
	}
	for _, n := range many {
		if b.Bit(n) == false {
			return false,nil
		}
	}
	return true,nil
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

// clears one bit
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
// 		not sure about best format - this seems excessively verbose
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

// func (b *BitField) CountOfFalseBits() int {
// }

// func (b *BitField) CountOfTrueBits() int {
// }

// Print the whole field to stdout
// 		not sure about best format - this seems excessively verbose
func (b *BitField) Dump() {
	b.DumpLoHi(0, b.maxbit)
}

// Returns a slice with indices of the bits with value = false
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
func (b *BitField) OrBitsByNdx(many []int) (bool,error) {
	if len(many) <= 0 {
		if verbose {fmt.Printf("Warning: OrBitsByNdx(emptySet) isn't meaningful\n")}
		return false, errNonEmptySliceRequired 
	}
	for _, n := range many {
		if b.Bit(n) == true {
			return true,nil
		}
	}
	return false,nil
}

// Set one bit 
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

// return a slice with indices of the bits with value = true
func (b *BitField) TrueBitsLoHi(lo, hi int) []int {
	rv := make([]int, 0, b.maxbit)
	for i := lo; i <= hi; i++ {
		if b.Bit(i) {
			rv = append(rv, i)
		}
	}
	return rv
}










