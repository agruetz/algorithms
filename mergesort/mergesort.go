package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func sort(arrayToCount []int) ([]int, int) {
	if len(arrayToCount) <= 1 {
		return arrayToCount, 0
	}
	mid := (len(arrayToCount) / 2)
	firstHalfSorted, invLeft := sort(arrayToCount[:mid])
	secondHalfSorted, invRight := sort(arrayToCount[mid:])
	finalSorted, invSplit := merge(firstHalfSorted, secondHalfSorted, mid)

	return finalSorted, (invLeft + invRight + invSplit)
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

	_, i := sort(a)
	fmt.Printf("Total Number of Inversions: %d\n", i)
}
