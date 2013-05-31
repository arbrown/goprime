package main

import (
		"fmt"
		)

func main() {
	candidateChan := make(chan int)
	max := 5
	go produceCandidates(max,candidateChan)
	for i:=0;i<max;i++{
		fmt.Println(<-candidateChan)
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




