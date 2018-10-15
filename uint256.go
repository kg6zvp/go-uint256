package uint256

import (
	"log"

	"github.com/kg6zvp/go-intdian"
)

// length of a Key in uint64
const LENGTH = 4

// number of bytes in a uint64, for convenience
const STEP_SIZE = 8

// number of bytes in a Uint256
const BYTE_LENGTH = STEP_SIZE * LENGTH

type Uint256 [LENGTH]uint64

// Equal checks equality between this Uint256 and another one
// params:
//     b: Uint256 to compare against
// returns: bool
func (a *Uint256) Equal(b Uint256) bool {
	comp := getComp()
	it := getIt()
	for i := getStart(); comp(i); i = it(i) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IsEmpty checks whether this Uint256 is empty
// returns: bool
func (a *Uint256) IsEmpty() bool {
	return a.Equal(EmptyUint256())
}

// LessThan checks if the given Uint256 is less than this Uint256
// params:
//     b: Uint256 to compare against
// returns: bool
func (a *Uint256) LessThan(b Uint256) bool {
	if b.IsEmpty() { //if it's empty, we don't care to compare
		return false
	}
	comp := getComp()
	it := getIt()
	for i := getStart(); comp(i); i = it(i) {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		} // continue to a less significant digit if they are equal
	}
	return false
}

// GreaterThan checks if the given Uint256 is more than this Uint256
// params:
//     b: Uint256 to compare against
// returns: bool
func (a *Uint256) GreaterThan(b Uint256) bool {
	return !a.Equal(b) && !a.LessThan(b)
}

// Xor calculates the xor of two Uint256's
// params:
//     b: Uint256 to compare against
// returns: Uint256 representing the difference
func (a *Uint256) Xor(b Uint256) Uint256 {
	var output Uint256
	output = Uint256{}
	comp := getComp()
	it := getIt()
	for i := getStart(); comp(i); i = it(i) {
		output[i] = a[i] ^ b[i]
	}
	return output
}

// ToBytes converts this Uint256 into a []byte
// returns: []byte
func (a *Uint256) ToBytes() []byte {
	out := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	comp := getComp()
	it := getIt()
	for i := getStart(); comp(i); i = it(i) {
		intdian.ByteOrder().PutUint64(out[i*STEP_SIZE:], a[i])
	}
	return out
}

// FACTORY METHODS

// FromBytes creates a Uint256 from a [32]byte
// params:
//     data: []byte
// returns: Uint256
// Fatal error for incorrect data length
func FromBytes(data []byte) Uint256 {
	if len(data) != BYTE_LENGTH {
		log.Fatalf("Incorrect data length: %d", len(data))
	}
	return Uint256{
		intdian.ByteOrder().Uint64(data[0:8]),
		intdian.ByteOrder().Uint64(data[8:16]),
		intdian.ByteOrder().Uint64(data[16:24]),
		intdian.ByteOrder().Uint64(data[24:32]),
	}
}

// EmptyUint256 returns an empty Uint256
// returns: Uint256
func EmptyUint256() Uint256 {
	return Uint256{0, 0, 0, 0}
}

// New returns a Uint256 with ordered most to least significant digits as params
// params:
//     num: ...uint64 most-to-least significant digit
// returns: Uint256
func New(num ...uint64) Uint256 {
	if len(num) > LENGTH {
		log.Fatalf("Wrong number of arguments given. Expected <=%d, got %d", LENGTH, len(num))
	}
	comp := getComp()
	it := getIt()
	v := EmptyUint256()

	// https://imgur.com/gallery/0QmLNmx
	for i, k := skipBy(it, getStart(), LENGTH-len(num)), 0; comp(i) && k < LENGTH; i, k = it(i), k+1 {
		v[i] = num[k]
	}
	return v
}
