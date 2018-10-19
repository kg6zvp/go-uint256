package uint256

import (
	"github.com/kg6zvp/go-intdian"
)

// most to least significant
func getStart() int {
	if intdian.Big_Endian {
		return 0
	} else {
		return LENGTH - 1
	}
}

// least to most significant
func getRevStart() int {
	if intdian.Big_Endian {
		return LENGTH - 1
	} else {
		return 0
	}
}

type cond func(i int) bool

type iter func(i int) int

// most to least significant
func getComp() cond {
	if intdian.Big_Endian {
		return func(i int) bool {
			return i < LENGTH
		}
	} else { // little endian
		return func(i int) bool {
			return i >= 0
		}
	}
}

// least to most significant
func getRevComp() cond {
	if intdian.Big_Endian {
		return func(i int) bool {
			return i >= 0
		}
	} else { // little endian
		return func(i int) bool {
			return i < LENGTH
		}
	}
}

// most to least significant
func getIt() iter {
	if intdian.Big_Endian {
		return func(i int) int {
			return i + 1
		}
	} else { // little endian
		return func(i int) int {
			return i - 1
		}
	}
}

// least to most significant
func getRevIt() iter {
	if intdian.Big_Endian {
		return func(i int) int {
			return i - 1
		}
	} else { // little endian
		return func(i int) int {
			return i + 1
		}
	}
}

// skip ahead by the given number of iterations from the given start
// params:
//     it: iter the iterator function
//     start: int the starting point to iterate from
//     num: the number of times to skip ahead by
// returns: int
func skipBy(it iter, start int, num int) int {
	for i := 0; i < num; i++ {
		start = it(start)
	}
	return start
}
