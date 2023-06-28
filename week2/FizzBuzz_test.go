package main

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for i := 1; i <= 100; i++ {
		expectedResult := getExpectedResult(i)
		if result := FizzBuzz(i); result != expectedResult {
			t.Errorf("Test failed for number %d: Expected %s, got %s", i, expectedResult, result)
		}
	}
}

func getExpectedResult(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}