package sorting

import "log"

func Sort(s []int) []int {
	iterationsCount := 0
	for i := range s {
		for j := i; j > 0; j-- {
			iterationsCount++
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			} else {
				break
			}
		}
	}
	log.Println("Iterations Count:", iterationsCount)
	return s
}


