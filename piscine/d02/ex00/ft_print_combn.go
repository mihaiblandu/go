package main

func prinComb(arr []int) {
	i := 0
	for i < len(arr) {
		print(i)
		i++
	}
	print(", ")
}

func ftPrintComb2(comb int) {
	var arr [comb]int
}
func init(arr []int) {
	i := 0
	for i < len(arr) {
		arr[0] = i
		i++
	}
	makeCombo(arr, len(arr)-1)
}
func makeCombo(arr []int, step int) {
	if(arr[string] < 9 ){
		arr[step]++
		makeCombo(arr []int, step int)
	}else if(step > 0){
		arr[step] =arr[ step -1] +1
	
	}
}
