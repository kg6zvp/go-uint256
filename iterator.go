package uint256

import (
	"github.com/kg6zvp/go-intdian"
)

func getStart() int {
	if intdian.Big_Endian {
		return 0
	} else {
		return LENGTH - 1
	}
}

type cond func(i int) bool

type iter func(i int) int

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
