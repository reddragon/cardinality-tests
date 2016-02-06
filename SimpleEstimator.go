package main

import (
  "fmt"
  "math"
  "math/rand"
)

type SimpleEstimator struct {
  maxZeroes uint
  rng *rand.Rand
}

func (b *SimpleEstimator) Init(seed int64) {
  b.maxZeroes = 0
  b.rng = rand.New(rand.NewSource(seed))
}

func (b *SimpleEstimator) newElement() {
  rInt := b.rng.Int31()
  rIntC := rInt

  var zeroes uint
  for zeroes = 0; rIntC % 2 == 0; rIntC /= 2 {
    zeroes++
    if zeroes > b.maxZeroes {
      b.maxZeroes = zeroes
    }
  }
}

func (b *SimpleEstimator) estimate() uint {
  return 1 << b.maxZeroes
}

func main() {
  estimator := new(SimpleEstimator)
  estimator.Init(42)

  for i := 0; i < 1024 * 1024; i++ {
    estimator.newElement()
    estimated := estimator.estimate()
    if i % 128 == 0 {
      // fmt.Printf("#: %d, Estimated: %d, Real: %d, Difference: %f\n",
      //   i, estimated, i + 1, math.Abs(1 - (float64(estimated) / float64(i + 1))))

      fmt.Printf("%d  %f\n",
        i, math.Abs(1 - (float64(estimated) / float64(i + 1))))
    }
  }
}
