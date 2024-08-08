package codility

import "testing"

func TestFindPosiNumSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{A: []int{-100, 3, 4}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPosiNumSolution(tt.args.A); got != tt.want {
				t.Errorf("FindPosiNumSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FindPosiNumSolution(A []int) int {
	// Implement your solution here
	lenth := len(A)
	result := make(map[int]bool, 100000)

	for i := 0; i < lenth; i++ {
		result[A[i]] = true
	}
	for j := 1; j <= 100000; j++ {
		_, ok := result[j]
		if !ok {
			return j
		}
	}
	return 100001
}
