/*
An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

Write a function:

func Solution(A []int, K int) []int

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given

    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given

    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]

Given

    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.


*/

package codility

import (
	"reflect"
	"testing"
)

func SolutionRotation(A []int, K int) []int {
	// Implement your solution here
	length := len(A)
	if length == 0 || K < 0 {
		return A
	}
	k := K % length
	space := length - k
	result := A[space:length]
	for i := 0; i < space; i++ {
		result = append(result, A[i])
	}
	return result
}

func TestSolutionRotation(t *testing.T) {

	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "empty", args: args{A: []int{}, K: 1}, want: []int{}},
		{name: "test", args: args{A: []int{1, 2, 3, 4}, K: 1}, want: []int{4, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolutionRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
