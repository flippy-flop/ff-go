package main

import (
	"math/rand"
	"fmt"
)

type FlippyFlop struct {
	data[] int
}

func (ff *FlippyFlop) Count() int {
	return len(ff.data)
}

func (ff *FlippyFlop) Flip(i int) bool {
	success := rand.Intn(2) == 1
	if success {
		if ff.Count() == cap(ff.data) {
			int[] growth = make(int[], 10)
			ff.data = append(ff.data, growth)
		}
		position := rand.Intn(ff.Count())
		ff.data[position] = i
	}
	return success
}

func (ff *FlippyFlop) Flop() (int, bool) {

	if ff.Count() > 0 {
		shouldRemove := rand.Intn(2) == 1
		position := rand.Intn(ff.Count())
		ret := ff.data[position]

		if shouldRemove {
			ff.data = append(ff.data[:position], ff.data[position+1:]...)
		}

		return ret, shouldRemove
	} else {
		return -1, false
	}
}

func main() {
	ff := FlippyFlop{}
	ff.Flip(5)
	fmt.Printf("Count: %v\n", ff.Count())
	ff.Flop()
	ff.Flop()
	fmt.Printf("Count: %v\n", ff.Count())
}