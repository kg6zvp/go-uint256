# Uint256

Unsigned 256-bit integer library for all your cryptographic needs

## Example

```go
import (
    "github.com/kg6zvp/go-uint256"
)

zero := uint256.Uint256{0,0,0,0} // a uint256.Uint256 is made up of 4 uint64's
if zero.Equal(uint256.EmptyUint256()) {
	//this will be true
}
```
