package common

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

// Mimax returns min and max from a list of numbers.
func Mimax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return min, max
}

// Random rertuns a random number over a range
func Random(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// ChanToSlice pushes values from a channel to a slice.
func ChanToSlice(ch chan int) []int {
	out := []int{}

	for v := range ch {
		out = append(out, v)
	}

	return out
}

// Contain checks if the target is in a slice.
func Contain(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

// ContainString checks if the target is in a slice.
func ContainString(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

// IsMoreThan1Apart checks if two integers are more than 1 apart.
func IsMoreThan1Apart(a, b int) bool {
	if math.Abs(float64(a)-float64(b)) > 1 {
		return true
	}

	return false
}

// IsLessThan1Apart checks if two integers are less or equal than 1 apart.
func IsLessThan1Apart(a, b int) bool {
	if math.Abs(float64(a)-float64(b)) <= 1 {
		return true
	}

	return false
}

// Log prints out the map of logging context and value.
func Log(m map[string]interface{}) {
	fmt.Println("[debug] →")
	for k, v := range m {
		fmt.Printf("\t%v: %+v\n", k, v)
	}
	fmt.Println("[debug] □")
}

// Equal checks if two input are deeply equal.
func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}
