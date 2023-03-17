package idiom

import (
	"math/big"
	"testing"
	"time"
)

func Test_Profiling(t *testing.T) {
	defer Duration(time.Now(), "Factorial")

	x := big.NewInt(1024)
	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}
}
