package main

import (
	"fmt"
	"sort"
)

// Activity is a [start, finish) interval.
type Activity struct {
	Start, Finish int
}

// Select returns a maximum-size set of non-overlapping activities, chosen
// by earliest finish time. Touching activities (one ends exactly when the
// next starts) are compatible — unlike the meetings variant.
func Select(activities []Activity) []Activity {
	sorted := append([]Activity{}, activities...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Finish < sorted[j].Finish })
	out := []Activity{}
	lastFinish := -1 << 62
	for _, a := range sorted {
		if a.Start >= lastFinish {
			out = append(out, a)
			lastFinish = a.Finish
		}
	}
	return out
}

func main() {
	fmt.Println(Select([]Activity{{10, 20}, {12, 25}, {20, 30}}))
}
