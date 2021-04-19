package _069_Sqrt_x

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMySqrt(t *testing.T) {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	numCount := 1000
	for i := 0; i < numCount; i++ {
		num := rnd.Intn(math.MaxInt32)
		fact, expected := mySqrt(num), mySqrt2(num)
		if fact != expected {
			t.Errorf("num: %d, want sqrt %d, got sqrt %d", num, expected, fact)
		}
	}
}
