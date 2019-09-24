package quicksort

import "fmt"

func Sort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	var i int
	for i = len(s)-1; i > 0; i-- {
		if s[i] < s[i-1] {
			s[i], s[i-1] = s[i-1], s[i]
		} else {
			break
		}
	}
	slice1, slice2 := s[:i], s[i:]
	fmt.Println(slice1, slice2)
	return slice1
}