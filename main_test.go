package main

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	cases := []struct {
		in, want []Activity
	}{
		{
			[]Activity{{10, 20}, {12, 25}, {20, 30}},
			[]Activity{{10, 20}, {20, 30}}, // touching intervals are compatible
		},
		{
			[]Activity{{1, 2}, {3, 4}, {0, 6}, {5, 7}, {8, 9}, {5, 9}},
			[]Activity{{1, 2}, {3, 4}, {5, 7}, {8, 9}},
		},
		{[]Activity{}, []Activity{}},
		{[]Activity{{1, 10}}, []Activity{{1, 10}}},
	}
	for _, c := range cases {
		if got := Select(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Select(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
