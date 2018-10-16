package uint256

import (
	"fmt"
	"testing"

	"github.com/kg6zvp/go-intdian"
)

func TestGetStart(t *testing.T) {
	if intdian.BigEndian() {
		if getStart() != 0 {
			t.Fatalf("Big endian iterators should start at 0, started at %d", getStart())
		}
	} else { // little endian
		if getStart() != 3 {
			t.Fatalf("Little endian iterators should start at 3, started at %d", getStart())
		}
	}
}

func TestCompare(t *testing.T) {
	comp := getComp()
	if intdian.Big_Endian {
		if comp(LENGTH + 1) { // shouldn't allow larger numbers
			t.Fail()
		}
		if !comp(-1) { //should allow low numbers
			t.Fail()
		}
	} else { // little endian
		if !comp(LENGTH + 1) { // should allow larger numbers
			t.Fail()
		}
		if comp(-1) { //shouldn't allow low numbers
			t.Fail()
		}
	}
}

func TestIterator(t *testing.T) {
	it := getIt()
	if intdian.Big_Endian {
		if it(1) != 2 {
			t.Fail()
		}
	} else { // little endian
		if it(1) != 0 {
			t.Fail()
		}
	}
}

func TestIteration(t *testing.T) {
	start := getStart()
	it := getIt()
	comp := getComp()
	if intdian.Big_Endian {
		bigIteration(t, start, comp, it)
	} else { // little endian
		littleIteration(t, start, comp, it)
	}
}

func bigIteration(t *testing.T, start int, comp cond, it iter) {
	if start != 0 {
		t.Fatalf("big endian should start at 0, started at %d", start)
	}

	ti := start
	for i := 0; i < 4; i++ {
		fmt.Printf("i = %d, ti = %d\n", i, ti)
		if ti > 3 || ti < 0 {
			t.Fatalf("ti should be between 0 and 3, inclusive, was %d", ti)
		}
		if !comp(ti) {
			t.Fatalf("comparison failed at ti == %d", ti)
		}
		ti = it(ti)
	}

	if ti != 4 {
		t.Fatalf("big endian should end at 4, ended at %d", ti)
	}
}

func littleIteration(t *testing.T, start int, comp cond, it iter) {
	if start != 3 {
		t.Fatalf("little endian should start at 3, started at %d", start)
	}

	ti := start
	for i := 3; i >= 0; i-- {
		fmt.Printf("i = %d, ti = %d\n", i, ti)
		if ti > 3 || ti < 0 {
			t.Fatalf("ti should be between 3 and 0, inclusive, was %d", ti)
		}
		if !comp(ti) {
			t.Fatalf("comparison failed at ti == %d", ti)
		}
		ti = it(ti)
	}

	if ti != -1 {
		t.Fatalf("little endian should end at -1, ended at %d", ti)
	}
}
