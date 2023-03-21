package main

import "fmt"

// Interface que define o comportamento dos algoritmos
type SortAlgorithm interface {
	Sort(data []int) []int
}

// Estrutura que representa o algoritmo bubble sort
type BubbleSort struct{}

// Método que implementa o algoritmo bubble sort
func (bs *BubbleSort) Sort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// Estrutura que representa o algoritmo quick sort
type QuickSort struct{}

// Método que implementa o algoritmo quick sort
func (qs *QuickSort) Sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	pivot := data[0]
	var left []int
	var right []int
	for _, num := range data[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = qs.Sort(left)
	right = qs.Sort(right)
	return append(append(left, pivot), right...)
}

// Estrutura que representa o objeto que utiliza os algoritmos
type Sorter struct {
	algorithm SortAlgorithm
}

// Método que ordena os dados utilizando o algoritmo definido
func (s *Sorter) Sort(data []int) []int {
	return s.algorithm.Sort(data)
}

func main() {
	// Uso do Strategy para selecionar o algoritmo de ordenação
	bubbleSort := &BubbleSort{}
	quickSort := &QuickSort{}

	sorter := &Sorter{algorithm: bubbleSort}
	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	sortedData := sorter.Sort(data)
	fmt.Println(sortedData) // [1 1 2 3 3 4 5 5 5 6 9]

	sorter.algorithm = quickSort
	sortedData = sorter.Sort(data)
	fmt.Println(sortedData) // [1 1 2 3 3 4 5 5 5 6 9]
}
