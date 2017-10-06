// Author: Garry Cummins
// Date: 06/10/2017

// Adapted from: https://play.golang.org/p/Ma2GXvj3XP,
//https://github.com/ZachOrr/golang-algorithms/blob/master/sorting/merge-sort.go
//https://play.golang.org/p/g3QTWcy9D-


package main

import (
	"fmt"
)
func Merge(l, r []int) []int {

	ret := make([]int, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}
//Function that is called to merge the two arrays
func MergeSort(s []int) []int {

	if len(s) <= 1 {
		return s
	}

	h := len(s) / 2
	l := MergeSort(s[:h])
	r := MergeSort(s[h:])
	return Merge(l, r)
}

func mergeArray(s1, s2 []int) []int {

	unique := make(map[int]struct{})

	for _, v := range s1 {
		unique[v] = struct{}{}
	}

	for _, v := range s2 {
		unique[v] = struct{}{} 
	}

	final := make([]int, len(unique)) 
	i := 0

	for k := range unique {
		final[i] = k
		i++ 
	}
	return final
}

func main() {
	// First Array
	f := []int{
		9,18,27,26,
		5,10,15,20,
	}
	//Second Array
	s := []int{
		7,65,12,40,
		1,70,55,32,
	}

	// Merging the two arrays together 
	merged := mergeArray(f, s)

	fmt.Printf(" First Array:  %v\n Second Array: %v\n Merged Array: %v\n", f, s, MergeSort(merged))
}