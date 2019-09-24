package main

import (
	"./sorting/quicksort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//slice := generateIntSlice(1000000)
	//timeStart := time.Now()
	//slice = sorting.Sort(slice)
	//fmt.Println("Execution Time: ", time.Since(timeStart).Seconds())

	//slice := generateIntSlice(10)
	fmt.Println(quicksort.Sort(&[]int{3,7,2,1,6,4}))
}

func generateIntSlice(l int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice = append(slice, rand.Intn(10000000000))
	}
	return slice
}