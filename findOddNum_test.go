/*
A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the elements at indexes 0 and 2 have value 9,
the elements at indexes 1 and 3 have value 3,
the elements at indexes 4 and 6 have value 9,
the element at index 5 has value 7 and is unpaired.
Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the function should return 7, as explained in the example above.

Write an efficient algorithm for the following assumptions:

N is an odd integer within the range [1..1,000,000];
each element of array A is an integer within the range [1..1,000,000,000];
all but one of the values in A occur an even number of times.
*/

package codility

import "testing"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func OddSolution(A []int) int {
	// Implement your solution here
	var x int
	result := make(map[int]int, 500000)
	for _, v := range A {
		_, ok := result[v]
		if ok {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}
	for k, v := range result {
		if v%2 == 1 {
			x = k
			break
		}
	}
	return x
}

func TestOddSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]int{1, 2, 3, 4, 7, 4, 3, 2, 1}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddSolution(tt.args.A); got != tt.want {
				t.Errorf("OddSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
