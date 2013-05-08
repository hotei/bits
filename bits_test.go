// BitField_tests.go

// NOTE: verbose is not public, but it's available to bits_test since it's the same pkg
//		tests should run ok even if verbose is false

// go test                   runs tests, but not benchmarks
// go test -test.bench=".*"  runs tests and benchmarks (use verbose=false for this)
package bits

import (
	"fmt"
	"os"
	"testing"
)

// ==========================================================   T E S T S

func Test_01(t *testing.T) {
	fmt.Printf("Test_01...\n")
	verbose = true
	var bits BitField
	bits.SetMaxBitNdx(10)
	bits.SetName("Test_01")
	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	bits.SetBit(0)
	if bits.Bit(0) != true {
		t.Fatalf("SetBit() failed")
	}
	if bits.Bit(1) == true {
		t.Fatalf("SetBit() failed")
	}

	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	bits.SetBit(9)
	if bits.Bit(9) != true {
		t.Fatalf("SetBit() failed")
	}
	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	bits.ClrBit(9)
	if bits.Bit(9) != false {
		t.Fatalf("ClrBit() failed")
	}
}

func Test_02(t *testing.T) {

	fmt.Printf("Test_02...\n")
	var bits BitField
	bits.SetMaxBitNdx(20)
	bits.SetName("Test_02")

	bits.SetBits([]int{4, 12})
	if bits.Bit(12) != true {
		t.Fatalf("SetBit() failed")
	}

	bits.TglBit(12)
	if bits.Bit(12) != false {
		t.Fatalf("TglBit() failed")
	}

	bits.TglBit(12)
	if bits.Bit(12) != true {
		t.Fatalf("TglBit() failed")
	}

	bits.ClrBit(12)
	if bits.Bit(12) != false {
		t.Fatalf("ClrBit() failed")
	}
}

func Test_03(t *testing.T) {
	fmt.Printf("Test_03...\n")
	var bits BitField
	bits.SetMaxBitNdx(10)
	bits.SetName("Test_03")

	var x = []int{0, 2, 4, 16}
	bits.SetBits(x)
	fmt.Printf("%s\n", bits.HexString())
	truebits := bits.TrueBitsLoHi(0, 16)
	fmt.Printf("should be {0, 2, 4, 16} %v\n", truebits)

	bits.ClrBits(x)
	fmt.Printf("%s\n", bits.HexString())
	truebits = bits.TrueBitsLoHi(0, 16)
	fmt.Printf("should be {} %v\n", truebits)
	if len(truebits) != 0 {
		t.Fatalf("ClrBits() failed")
	}
}

func Test_04(t *testing.T) {
	fmt.Printf("\nTest_04...\n")
	var bits BitField
	bits.SetMaxBitNdx(10)
	bits.SetName("Test_04")

	var x = []int{0, 2, 4, 16}
	bits.SetBits(x)
	tStr := "a80080"
	fmt.Printf("%s\n", bits.HexString())
	if bits.HexString() != tStr {
		t.Fatalf("bits.HexString() failed")
	}
	tStr = "101010000000000010000000"
	if bits.String() != tStr {
		t.Fatalf("bits.String() failed")
	}
	fmt.Printf("This ruler helps to count bits\n")
	fmt.Printf("          1         2         3\n")
	fmt.Printf("0123456789012345678901234567890\n")
	fmt.Printf("%s\n", bits.String())
	bits.DumpLoHi(2, 4)
}

func Test_05(t *testing.T) {
	fmt.Printf("\nTest_05...\n")
	var bits BitField
	bits.SetMaxBitNdx(10)
	bits.SetName("Test_05")

	var x = []int{0, 2, 4, 16}
	bits.SetBits(x)
	//bits.Dump()
	//bits.DumpLoHi(2,4)
	tv, err := bits.OrBitsByNdx([]int{3, 7, 5})
	if err != nil {
		fmt.Printf("OrBitsByNdx() %v\n", err)
		os.Exit(-1)
	}
	if tv != false {
		t.Fatalf("OrBitsByNdx() failed")
	}

	tv, err = bits.OrBitsByNdx([]int{3, 7, 5, 2})
	if err != nil {
		fmt.Printf("OrBitsByNdx() %v\n", err)
		os.Exit(-1)
	}
	if tv != true {
		t.Fatalf("OrBitsByNdx() failed")
	}

	tv, err = bits.AndBitsByNdx([]int{1, 2, 3})
	if err != nil {
		fmt.Printf("OrBitsByNdx() %v\n", err)
		os.Exit(-1)
	}
	if tv != false {
		t.Fatalf("OrBitsByNdx() failed")
	}

	tv, err = bits.AndBitsByNdx([]int{2, 4, 0})
	if err != nil {
		fmt.Printf("OrBitsByNdx() %v\n", err)
		os.Exit(-1)
	}
	if tv != true {
		t.Fatalf("OrBitsByNdx() failed")
	}

}

func Test_06(t *testing.T) {
	fmt.Printf("Test_06...\n")
	var bits BitField
	bits.SetMaxBitNdx(31)
	bits.SetName("Test_06")

	bits.SetLoHi(0, 31)
	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	for i := 0; i <= 31; i++ {
		if bits.Bit(i) != true {
			t.Fatalf("SetLoHi() failed at [%d]", i)
		}
	}
	bits.ClrLoHi(0, 31)
	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	for i := 0; i <= 31; i++ {
		if bits.Bit(i) != false {
			t.Fatalf("SetLoHi() failed at [%d]", i)
		}
	}
	bits.TglLoHi(0, 31)
	if verbose {
		fmt.Printf("%s\n", bits.HexString())
	}
	for i := 0; i <= 31; i++ {
		if bits.Bit(i) != true {
			t.Fatalf("SetLoHi() failed at [%d]", i)
		}
	}
}

// test self-extending feature
func Test_07(t *testing.T) {
	fmt.Printf("Test_07...\n")
	fmt.Printf("Self extending feature test\n")
	var bits BitField
	//bits.SetMaxBitNdx(16)  // left out on purpose, should not change result
	bits.SetName("Test_07")
	oldverbose := verbose
	verbose = false
	for testval := 0; testval < 100; testval++ {
		bits.SetBit(testval)
		s := bits.HexString()
		if verbose {
			fmt.Printf("%s\n", s)
		}
		units := (testval + 1) / 8
		if ((testval + 1) % 8) != 0 {
			units++
		}
		if verbose {
			fmt.Printf("Self extending feature implies bit[%d] should require %d units of storage (%d hex chars) %d\n",
				testval, units, units*2, len(s))
		}
		if len(s) != (units * 2) {
			t.Fatalf("Self extending feature failed bit[%d] should require %d units of storage (%d hex chars)\n",
				testval, units, units*2)
		}
	}
	verbose = oldverbose
}

// ==========================================================   B E N C H M A R K S

func BenchmarkBitSet(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(15)

	for i := 0; i < b.N; i++ {
		bits.SetBit(2)
	}
}

func BenchmarkBitClr(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(15)

	for i := 0; i < b.N; i++ {
		bits.ClrBit(2)
	}
}

func BenchmarkBitTgl(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(15)

	for i := 0; i < b.N; i++ {
		bits.TglBit(2)
	}
}

func BenchmarkBitRead(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(15)

	for i := 0; i < b.N; i++ {
		_ = bits.Bit(2)
	}
}

func BenchmarkBitSetMany(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(15)

	var x = []int{0, 2, 4, 16}
	for i := 0; i < b.N; i++ {
		bits.SetBits(x)
	}
}

func BenchmarkBitClrMany(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(16)

	var x = []int{0, 2, 4, 16}
	for i := 0; i < b.N; i++ {
		bits.ClrBits(x)
	}
}

func BenchmarkBitTglMany(b *testing.B) {
	var bits BitField
	verbose = false
	bits.SetMaxBitNdx(16)

	var x = []int{0, 2, 4, 16}
	for i := 0; i < b.N; i++ {
		bits.TglBits(x)
	}
}
