package uint256

import (
	"reflect"
	"testing"

	"github.com/kg6zvp/go-intdian"
)

func TestFromBytes(t *testing.T) {
	var v Uint256
	v = FromBytes([]byte{43, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 2, 0, 0, 0, 0, 0})
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
	v := New(845)
	bs := v.ToBytes()
	if intdian.Big_Endian {
		if bs[30] != 3 {
			t.Fatalf("bs[30] should be 3, was %d", bs[30])
		}
		if bs[31] != 77 {
			t.Fatalf("bs[31] should be 77, was %d", bs[31])
		}
	} else { // little endian
		if bs[1] != 3 {
			t.Fatalf("bs[1] should be 3, was %d", bs[1])
		}
		if bs[0] != 77 {
			t.Fatalf("bs[0] should be 77, was %d", bs[0])
		}
	}
}

func TestConstructor(t *testing.T) {
	v := New()
	if !v.IsEmpty() {
		t.Fatalf("New() without args should be 0/empty, was {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
	}

	v = New(20)
	if intdian.Big_Endian {
		if !reflect.DeepEqual(v, Uint256{0, 0, 0, 20}) {
			t.Fatalf("big endian expected: {0, 0, 0, 20}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	} else { // little endian
		if !reflect.DeepEqual(v, Uint256{20, 0, 0, 0}) {
			t.Fatalf("little endian expected: {20, 0, 0, 0}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	}

	v = New(255, 36)
	if intdian.Big_Endian {
		if !reflect.DeepEqual(v, Uint256{0, 0, 255, 36}) {
			t.Fatalf("big endian expected: {0, 0, 255, 36}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	} else { // little endian
		if !reflect.DeepEqual(v, Uint256{36, 255, 0, 0}) {
			t.Fatalf("little endian expected: {36, 255, 0, 0}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	}

	v = New(96, 83, 67)
	if intdian.Big_Endian {
		if !reflect.DeepEqual(v, Uint256{0, 96, 83, 67}) {
			t.Fatalf("big endian expected: {0, 96, 83, 67}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	} else {
		if !reflect.DeepEqual(v, Uint256{67, 83, 96, 0}) {
			t.Fatalf("little endian expected: {67, 83, 96, 0}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	}

	v = New(8196, 96, 83, 67)
	if intdian.Big_Endian {
		if !reflect.DeepEqual(v, Uint256{8196, 96, 83, 67}) {
			t.Fatalf("big endian expected: {8196, 96, 83, 67}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	} else {
		if !reflect.DeepEqual(v, Uint256{67, 83, 96, 8196}) {
			t.Fatalf("little endian expected: {67, 83, 96, 8196}, actual: {%d, %d, %d, %d}", v[0], v[1], v[2], v[3])
		}
	}
}
