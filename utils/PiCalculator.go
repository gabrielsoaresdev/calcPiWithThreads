package utils

import (
	"math"
	"time"
)

type piCalculator struct {
	termsQnt int
	threadsQnt int
}

func NewPiCalculator (termsQnt, threadsQnt int) *piCalculator {
	return &piCalculator{
		termsQnt:   termsQnt,
		threadsQnt: threadsQnt,
	}
}

func (pc piCalculator) run(c chan float64, initialN, termsQnt int) {
	var pi float64

	for n := initialN; n < initialN +termsQnt; n++ {
		pi += math.Pow(-1, float64(n)) / float64(2 * n + 1)
	}

	pi *= 4
	c <- pi
}

// Calculate function receives the number of terms and threads,
// and returns the value of PI and the Duration object, with
// the spent time in the operation.
func (pc piCalculator) Calculate() *ResultSet {
	start := time.Now()

	termsQuantityByThread := pc.termsQnt / pc.threadsQnt
	initialNumber := 0

	var channels []chan float64
	for j := 0; j < pc.threadsQnt; j++ {
		channels = append(channels, make(chan float64))
		go pc.run(channels[j], initialNumber, termsQuantityByThread)

		initialNumber += termsQuantityByThread
	}

	var pi float64
	for _, c := range channels {
		pi += <-c
	}

	spentTime := time.Since(start)
	return &ResultSet{pi, spentTime}
}