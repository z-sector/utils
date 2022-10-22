package utils

func InSlice(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}

func InSliceInt(s []int, x int) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}
