package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}
func IntSliceToString(a []int) string {
	var res string
	for _, el := range a {
		res += strconv.Itoa(el)
	}
	return res
}

func MergeSlices(f []float32, i []int32) []int {
	var res []int
	for _, el := range f {
		res = append(res, int(el))
	}
	for _, el := range i {
		res = append(res, int(el))
	}
	return res
}
func GetMapValuesSortedByKey(m map[int]string) []string {
	var res []string
	var index []int
	for key, _ := range m {
		index = append(index, key)
	}
	sort.Ints(index)
	for _, el := range index {
		res = append(res, m[el])
	}
	return res
}

func main() {}
