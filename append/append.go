package append

func Append(s []int, i int) []int {
	if len(s)+1 > cap(s) {
		// make a anew slice and copy everything from the old one
		newSlice := make([]int, len(s)+1, (len(s)+1)*2)
		for i, v := range s {
			newSlice[i] = v
		}
		newSlice[len(s)] = i
		s = newSlice
	} else {
		s = s[:len(s)+1]
		s[len(s)-1] = i
	}
	return s
}
