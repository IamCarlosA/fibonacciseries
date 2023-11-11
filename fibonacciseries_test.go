package fibonacciseries

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"zero elements", 0, []int{}},
		{"one elements", 1, []int{0}},
		{"two elements", 2, []int{0, 1}},
		{"five elements", 5, []int{0, 1, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate(%d) = %v, must %v", tt.n, got, tt.want)
			}
		})
	}
}
