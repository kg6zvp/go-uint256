package uint256

import (
	"testing"
)

func TestFromBytes(t *testing.T) {
	var v Uint256
	v = FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 43, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 8})
	if uint64(43) != v[0] {
		t.Fatalf("First 64bit int should've been 43, was %d", v[0])
	}
	if uint64(131080) != v[3] {
		t.Fatalf("Last 64bit int should've been 8, was %d", v[3])
	}
}

func TestEmptyUint256(t *testing.T) {
	var v Uint256
	v = EmptyUint256()
	for i, v := range v {
		if v != uint64(0) {
			t.Fatalf("%d int was non-zero: %d", i+1, v)
		}
	}
}

func TestToBytes(t *testing.T) {
	v := Uint256{0, 0, 0, 845}
	if v[3] != uint64(845) {
		t.Fail()
	}
	if v.ToBytes()[30] != 3 {
		t.Fail()
	}
	if v.ToBytes()[31] != 77 {
		t.Fail()
	}
}

func BenchmarkFromBytes(b *testing.B) {

	testBytes := []byte{0, 0, 0, 0, 0, 0, 0, 43, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 6, 0, 0, 0, 0, 0, 2, 0, 8}

	for n := 0; n < b.N; n++ {
		FromBytes(testBytes)
	}
}

func BenchmarkToBytes(b *testing.B) {

	//Init byte array before the benchmark
	var testBytes = make([]byte, 32)

	var v Uint256

	for n := 0; n < b.N; n++ {
		//Write to a variable to avoid inlining
		testBytes = v.ToBytes()
	}

	//Discard the value
	_ = testBytes

}
