package channel

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_printSequence(t *testing.T) {
	tests := []struct {
		n int
	}{
		// {1},
		// {2},
		// {3},
		// {4},
		// {5},
		{6},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			printSequence(tt.n)
		})
	}
}

func Test_printNumberLetter(t *testing.T) {
	printNumberLetter()
}

func Test_printOrder(t *testing.T) {
	tests := []struct {
		count int
	}{
		{3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%03d", tt.count), func(t *testing.T) {
			printOrder(tt.count)
		})
	}
}
