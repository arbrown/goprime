package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var primesToPrint int
	// Figure out how many primes to print
	if len(os.Args) > 1 {
		primesToPrint, _ = strconv.Atoi(os.Args[1])
	} else {
		primesToPrint = 5
	}

	printChan := make(chan int)

	// Produce all primes, sending them to printChan
	Primes(printChan)

	// Print as many primes as you like
	for i := 0; i < primesToPrint; i++ {
		fmt.Println(<-printChan)
	}
}

// Produce primes indefinitely
func Primes(outChan chan int) {
	candidateChan := make(chan int)

	// Produce potential prime numbers by weeding out
	// multiples of 2 and 3 (after 2 and 3, of course)
	go produceCandidates(candidateChan)

	// Start the filter at 0.  0 is never a candidate,
	// but a check for 0 is quick, and is a simple way to
	// guarantee we never skip an early candidate
	go (filter(0, candidateChan, outChan))()

}

// Produce prime number candidates, skipping multiples of 2 or 3
// after 3
func produceCandidates(ch chan int) {
	// Candidate producing algorithm starts
	// at 7
	cand := 7

	// Step over 4 or 2 integers alternatingly
	// to get odd numbers that are not multiples of 3
	gap := 4

	// Trivial prime numbers
	ch <- 2
	ch <- 3
	ch <- 5
	for {
		ch <- cand
		cand += gap
		gap = 6 - gap
	}
}

// Produce an anonymous function that will exclude
// multiples of each prime 'seed' from being prime
func filter(seed int, in chan int, out chan int) func() {
	myout := out
	// The 0 filter just passes the test number on
	// until we pass 5, when normal filtering begins
	if seed == 0 {
		return func() {
			for i := range in {
				myout <- i
				if i == 5 {
					ch := make(chan int)
					myout = ch
					go (filter(i, ch, out))()
				}
			}
		}
	}

	test := seed
	gap := 4

	return func() {
		// Keep listening on 'in' as long as someone is providing
		for cand := range in {
			if cand != test {
				// Pass it to the next filter, or to 'out' if this
				// filter is the last
				myout <- cand
				// If this is the last filter, launch a new goroutine
				// at the end of the chain
				if myout == out {
					ch := make(chan int)
					myout = ch
					go (filter(cand, ch, out))()
				}
			}

			// If the candidate number is greater than this filter,
			// bump this filter up to the next multiple of the
			// seed that will be tested (skipping multiples of 2 and 3)
			if cand > test {
				test += gap * seed
				gap = 6 - gap
			}
		}
	}
}
