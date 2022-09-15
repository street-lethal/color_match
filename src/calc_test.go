package src

import (
	"reflect"
	"testing"
)

func Test_getNext(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"first", args{[]int{5, 4, 3, 2, 1}}, []int{4, 5, 3, 2, 1}},
		{"regular", args{[]int{1, 3, 2, 5, 4}}, []int{2, 1, 3, 5, 4}},
		{"regular", args{[]int{1, 2, 4, 5, 3}}, []int{5, 3, 2, 1, 4}},
		{"last", args{[]int{1, 2, 3, 4, 5}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
