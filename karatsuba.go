package main

import (
	"fmt"
	"math"
	"math/big"
)

func getNumberDigitsInt64(num int64) uint8 {
	var numDigits uint8

	if num == 0 {
		return 1
	}

	for num > 0 {
		numDigits++
		num = num / 10
	}

	return numDigits
}

func numberSplitInt64(num int64, numDigits uint8) (int64, int64) {
	divisor := int64(math.Pow(10, float64(numDigits)))

	if num >= divisor {
		return num / divisor, num % divisor
	}

	return 0, num
}

func karatsubaInt64(x, y int64) int64 {
	if x == 0 || y == 0 {
		return 0
	}

	xNumDigits := getNumberDigitsInt64(x)
	yNumDigits := getNumberDigitsInt64(y)

	if xNumDigits <= 2 || yNumDigits <= 2 {
		return x * y
	}

	a, b := numberSplitInt64(x, (xNumDigits / 2))
	c, d := numberSplitInt64(y, (yNumDigits / 2))

	s1 := karatsubaInt64(a, c)
	s2 := karatsubaInt64(b, d)
	//ac+ad+bc+bd
	s3 := karatsubaInt64(a, c) + karatsubaInt64(a, d) + karatsubaInt64(b, c) + karatsubaInt64(b, d)

	s4 := (s3 - s1) - s2

	s15 := s1 * int64(math.Pow(10, 4))
	s45 := s4 * int64(math.Pow(10, 2))

	return s15 + s2 + s45
}

func main() {
	//The first part shows how to use the standard library math/big to do arbitrary persision number
	//multiplication. Under the hood this library takes care of creating the appropriate sized data structures
	//preventing you from having to convert a string to smaller numbers that can be multiplied.
	//
	//Additionally it implements a number of optimized methods such as karatsuba multiplication and a number
	//of bitshift operations to gain the best performance in multiplying large numbers
	//
	// 8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184 - correct
	x, _, _ := big.ParseFloat("3141592653589793238462643383279502884197169399375105820974944592", 10, 256, big.ToZero)
	y, _, _ := big.ParseFloat("2718281828459045235360287471352662497757247093699959574966967627", 10, 256, big.ToZero)
	z, _, _ := big.ParseFloat("0", 10, 512, big.ToZero)
	answer := z.Mul(x, y)
	fmt.Printf("The big number multipled answer is:\n%f\n", answer)

	//This is an implementation of the karatsuba multiplication algorithm with out any of the need
	//to do any of the work to handle extremely large numbers. This somewhat defeats the purpose of
	//the karatsuba algorithm as it mainly shows its performance gains in large numbers. However,
	//this was implemented to show how it works and for learning purposes.
	fmt.Printf("Number is: %d\n", karatsubaInt64(5678, 1234))
}
