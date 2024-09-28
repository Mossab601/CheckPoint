package piscine

func ConcatAlternate(slice1,slice2 []int) []int {
	if len(slice1) < len(slice2){
		slice1,slice2 =slice2,slice1
	}
	result:=[]int{}
	for i, l := range slice1 {
		result = append(result, l)
		if i < len(slice2) {
			result = append(result, slice2[i])
		}
	}
	return result
}