
<center>
package bits
============
</center>

License details are at the end of this document. 
This document is (c) 2013 David Rook.

Comments can be sent to <hotei1352@gmail.com>

Description
-----------
This [package][1] creates and manipulates arbitrary length bitfields.

---

Installation
------------

```
go get github.com/hotei/bits
```

---

Take A Quick Peek Under The Hood
--------------------------------
```

import "github.com/hotei/bits"

func tryit() {
	// unknown size for setA 
	var setA bits.BitField
	// we know the size of this one
	var setB bits.BitField
	setB.SetMaxBitNdx(100)
	setB.SetBit(2)
	setB.Dump()
}

See test files for more examples.
```
Features
--------

Supported functions:  GoDoc.Org http://godoc.org/github.com/hotei/bits

* func (b *BitField) AndBitsByNdx(many []int) (bool, error)
	* returns and'ing of all the many bits selected by the slice indices
	* it's an error to apply this to an empty slice

* func (b *BitField) Bit(n int) bool
	* returns value of the one bit at the nth location

* func (b *BitField) ClrBit(n int)
	* clears one bit

* func (b *BitField) ClrBits(many []int)
	* Clear every bit position in the slice of integers

* func (b *BitField) ClrLoHi(lo, hi int)
	* Clear bits in range `[lo..hi]` (inclusive)

* func (b *BitField) Dump()
	* Print the whole field to stdout 

* func (b *BitField) DumpLoHi(lo, hi int)
	* Print bits in range `[lo..hi]` (inclusive) to stdout 

* func (b *BitField) FalseBitsLoHi(lo, hi int) []int
	* Returns a slice with indices of the bits with value = false

* func (b *BitField) HexDump() string
	* return hex representation of bitfield as string

* func (b *BitField) OrBitsByNdx(many []int) (bool,error)
	* returns or'ing of all the bits selected by the slice
	* it's an error to apply this to an empty slice

* func (b *BitField) SetBit(n int)
	* Set one bit
	
* func (b *BitField) SetBits(many []int)
	* Set every bit indicated by indices in slice
	
* func (b *BitField) SetLoHi(lo, hi int)
	* Set every bit in the range `[lo..hi]` (inclusive)

* func (b *BitField) SetMaxBitNdx(n int)
	* Hint to allow faster startup, bitfield will grow if needed
	
* func (b *BitField) SetName(newname string)
	* convenience name to attach to dump
	
* func (b *BitField) String() string
	* returns string with "010110" representation of bitfield
	* length is right zero filled to multiple of 8
	
* func (b *BitField) TglBit(n int)
	* Toggle bit at position n

* func (b *BitField) TglBits(many []int)
	* Toggle every bit selected by slice indices
	
* func (b *BitField) TglLoHi(lo, hi int)
	* Toggle every bit in the range [lo..hi]
	
* func (b *BitField) TrueBitsLoHi(lo, hi int) []int
	* return a slice with indices of the bits with value = true

<font color=red>
TODO
----

* func (b *BitField) CountOfTrueBits() int
	* fast access to true bit count with integer lookup - not a loop
	* not done yet - need to benchmark vs original

* func (b *BitField) CountOfFalseBits() int
	* fast access to false bit count with integer lookup - not a loop
	* not done yet

</font>

---

BENCHMARK
---------
```	
	Go 1.1rc2 on 4Ghz AMD/64
	BenchmarkBitSet-4	500000000	         5.69 ns/op
	BenchmarkBitClr-4	500000000	         5.89 ns/op
	BenchmarkBitTgl-4	500000000	         5.64 ns/op
	BenchmarkBitRead-4	500000000	         5.90 ns/op
	BenchmarkBitSetMany-4	100000000	    28.50 ns/op	 ( sets 4 different bits )
	BenchmarkBitClrMany-4	100000000	    29.50 ns/op  ( clrs 4 different bits )
	BenchmarkBitTglMany-4	100000000	    28.40 ns/op  ( tgls 4 different bits )
```

---
[1]: http://github.com/hotei/bits "github.com/hotei/bits"

License
-------
The 'bits' go package and demo programs are distributed under the Simplified BSD License:

> Copyright (c) 2013 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// EOF README.md  (this is a markdown document and tested OK with blackfriday)
