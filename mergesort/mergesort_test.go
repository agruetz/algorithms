package main

import (
	"bufio"
	//"fmt"
	"os"
	"runtime"
	"strconv"
	"testing"
)

var c []int
var result int

func init() {
	for i := 0; i < 5000000; i++ {
		c = append(c, i)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestSort(t *testing.T) {
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

	if i != 2407905288 {
		t.Errorf("Incorrect inversion count of %d should be %d", i, 2407905288)
	}
}

func TestSortMulti(t *testing.T) {
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

	_, i := sortMulti(a)

	if i != 2407905288 {
		t.Errorf("Incorrect inversion count of %d should be %d", i, 2407905288)
	}
}
func BenchmarkSortSmallFile(b *testing.B) {
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

	result = i
}

func BenchmarkSortSmallFileMulti(b *testing.B) {
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

	_, i := sortMulti(a)

	result = i
}

func BenchmarkSortLarge(b *testing.B) {
	var i int
	for n := 0; n < b.N; n++ {
		_, i = sortArray(c)
	}

	result = i
}

func BenchmarkSortLargeMulti(b *testing.B) {
	var i int
	for n := 0; n < b.N; n++ {
		_, i = sortMulti(c)
	}

	result = i
}

func BenchmarkSortLargeStdLib(b *testing.B) {
	var i int
	i = 1
	for n := 0; n < b.N; n++ {
		goSort(c)
	}

	result = i
}
