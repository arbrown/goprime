package main

import (
		"fmt"
		)

func main() {
	candidateChan := make(chan int)
	printChan := make(chan int)
	max := 25
	go produceCandidates(max,candidateChan)
	go (filter(5, candidateChan, printChan))()
	
	
	for i:=0;i<6;i++{
		fmt.Println(<-printChan)
	}
	
}

func produceCandidates(max int, ch chan int) {
	gap := 4
	cand := 7
	for i:=0;i<max;i++ {
		ch <- cand
		cand += gap
		gap = 6 - gap
	}
}

func filter(seed int, in chan int, out chan int) func() {
	test := seed
	gap := 4
	return func () {
		for {
			cand := <- in
			if cand >= test {
				test += gap * test
				gap = 6 - gap
			} else {
				out <- cand
			}
		}
	}
}




