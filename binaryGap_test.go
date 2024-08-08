package codility

import "testing"

/*
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

class Solution { public int solution(int N); }

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
*/

// 1041 15 32

func Solution(N int) int {
	// Implement your solution here
	var flag, count, n, y int
	flag = 0
	count = 0
	n = N
	space := 0
	for n != 0 {
		y = n % 2
		n = n / 2
		if y == 1 {
			if flag == 1 {
				if space > count {
					count = space
				}
				space = 0
			}
			flag = 1
		} else {
			if flag == 1 {
				space++
			}
		}
	}
	return count

}

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1041", args: args{N: 1041}, want: 5},
		{name: "15", args: args{N: 15}, want: 0},
		{name: "32", args: args{N: 32}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
