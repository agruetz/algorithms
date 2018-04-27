package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

var limit = make(chan int, runtime.NumCPU())

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func merge(left, right []int, mid int) ([]int, int) {
	arrayToReturn := make([]int, (len(left) + len(right)))

	l, r, d, splitInv := 0, 0, 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			arrayToReturn[d] = left[l]
			l++
		} else {
			arrayToReturn[d] = right[r]
			r++
			splitInv += mid - l
		}
		d++
	}

	for l < len(left) {
		arrayToReturn[d] = left[l]
		l++
		d++
	}

	for r < len(right) {
		arrayToReturn[d] = right[r]
		r++
		d++
	}

	return arrayToReturn, splitInv
}

func sortArray(arrayToCount []int) ([]int, int) {
	if len(arrayToCount) <= 1 {
		return arrayToCount, 0
	}
	mid := (len(arrayToCount) / 2)
	firstHalfSorted, invLeft := sortArray(arrayToCount[:mid])
	secondHalfSorted, invRight := sortArray(arrayToCount[mid:])
	finalSorted, invSplit := merge(firstHalfSorted, secondHalfSorted, mid)

	return finalSorted, (invLeft + invRight + invSplit)
}

func sortMulti(arrayToCount []int) ([]int, int) {
	if len(arrayToCount) <= 1 {
		return arrayToCount, 0
	}
	mid := (len(arrayToCount) / 2)

	var firstHalfSorted, secondHalfSorted []int
	var invLeft, invRight int
	var wg sync.WaitGroup
	wg.Add(2)

	select {
	case limit <- 1:
		go func() {
			firstHalfSorted, invLeft = sortMulti(arrayToCount[:mid])
			<-limit
			wg.Done()
		}()
	default:
		firstHalfSorted, invLeft = sortArray(arrayToCount[:mid])
		wg.Done()
	}

	select {
	case limit <- 1:
		go func() {
			secondHalfSorted, invRight = sortMulti(arrayToCount[mid:])
			<-limit
			wg.Done()
		}()
	default:
		secondHalfSorted, invRight = sortArray(arrayToCount[mid:])
		wg.Done()
	}

	wg.Wait()

	finalSorted, invSplit := merge(firstHalfSorted, secondHalfSorted, mid)

	return finalSorted, (invLeft + invRight + invSplit)
}

func goSort(arrayToSort []int) {
	sort.Ints(arrayToSort)
}

func main() {

	file, err := os.Open("lotsofnumbers")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	a := make([]int, 0)
	for scanner.Scan() {
		int, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		a = append(a, int)
	}

	_, i := sortArray(a)
	fmt.Printf("Total Number of Inversions: %d\n", i)

	_, i = sortMulti(a)
	fmt.Printf("Total Number of Inversions: %d\n", i)
}
