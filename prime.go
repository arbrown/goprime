package main

import (
		"fmt"
		"os"
		"strconv"
		)

func main() {
	var primesToPrint int
	if len(os.Args) > 1 {
		primesToPrint, _ = strconv.Atoi(os.Args[1])
	}else {
		primesToPrint = 5
	}	
	
	candidateChan := make(chan int)
	printChan := make(chan int)
	go produceCandidates(candidateChan)
	go (filter(5, candidateChan, printChan))()
	
	// 2,3,5 are trivial primes...
	
	fmt.Printf("2\n3\n5\n")
	for i:=3; i<primesToPrint; i++{
		fmt.Println(<-printChan)
	}
	
}

func produceCandidates(ch chan int) {
	gap := 4
	cand := 7
	for {
		ch <- cand
		cand += gap
		gap = 6 - gap
	}
	close(ch)
}

func filter(seed int, in chan int, out chan int) func() {
	test := seed
	gap := 4
	myout := out
	return func () {
	//fmt.Printf("Starting new goroutine with %d\n", seed)
		for cand:= range in {
			if cand == test {
			//fmt.Printf("%d == %d. And is a multiple of %d\n", cand, test, seed)
			} else {
				myout <- cand
				// If we just found a prime, launch a new goroutine with a new filter
				if  myout == out {
					//fmt.Printf("** Just printed prime (%d).  Launching new goroutine.\n", cand)
					ch := make(chan int)
					myout = ch
					go (filter(cand, ch, out))()
				}
			}
			if cand > test {
				//fmt.Printf("%d is greater than the test (%d).  Adding %d * %d\n", cand, test, seed, gap)
				test += gap * seed
				gap = 6 - gap
			}
		}
		close(out)
	}
}




