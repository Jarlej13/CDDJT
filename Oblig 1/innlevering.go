package main

func main() {
}

// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				swap(list[j], list[j+1])
	}
}

		func
		swap(i int, j int) {
		temp := i
		i = j
		j = temp
		}