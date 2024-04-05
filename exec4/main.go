package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variavel := rand.Perm(100)

	//fmt.Println("vetor desordenado: ", variavel, len(variavel))

	fmt.Printf("Inicio do InsertionSort 100 items: %v \n", time.Now())
	insertionSort(variavel)
	fmt.Printf("Fim do InsertionSort 100 items: %v \n", time.Now())

	fmt.Printf("Inicio do CombSort 100 items: %v \n", time.Now())
	combsort(variavel)
	fmt.Printf("Fim do CombSort 100 items: %v \n", time.Now())

	fmt.Printf("Inicio do Selection 100 items: %v \n", time.Now())
	Selection_Sort(variavel)
	fmt.Printf("Fim do Selection 100 items: %v \n", time.Now())

	variavel2 := rand.Perm(1000)
	fmt.Printf("Inicio do InsertionSort 1000 items: %v \n", time.Now())
	insertionSort(variavel2)
	fmt.Printf("Fim do InsertionSort 1000 items: %v \n", time.Now())

	fmt.Printf("Inicio do CombSort 1000 items: %v \n", time.Now())
	combsort(variavel2)
	fmt.Printf("Fim do CombSort 1000 items: %v \n", time.Now())

	fmt.Printf("Inicio do Selection 1000 items: %v \n", time.Now())
	Selection_Sort(variavel2)
	fmt.Printf("Fim do Selection 1000 items: %v \n", time.Now())

	variavel3 := rand.Perm(10000)
	fmt.Printf("Inicio do InsertionSort 10000 items: %v \n", time.Now())
	insertionSort(variavel3)
	fmt.Printf("Fim do InsertionSort 10000 items: %v \n", time.Now())

	fmt.Printf("Inicio do CombSort 10000 items: %v \n", time.Now())
	combsort(variavel3)
	fmt.Printf("Fim do CombSort 10000 items: %v \n", time.Now())

	fmt.Printf("Inicio do Selection 10000 items: %v \n", time.Now())
	Selection_Sort(variavel3)
	fmt.Printf("Fim do Selection 10000 items: %v \n", time.Now())

	//fmt.Println("\n vetor ordenado: ", variavel, len(variavel))

}

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	//return arr
}
func combsort(items []int) {
	var (
		n       = len(items)
		gap     = len(items)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}

func Selection_Sort(array []int) {
	var n = len(array)
	var min_index int
	var temp int
	for i := 0; i < n-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < n; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	//return array
}
