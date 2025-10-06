package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 {
		return []int{}
	}
	if size <= 0 {
		return nil
	}
	randSlice := make([]int, size)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range size {
		randSlice[i] = rnd.Int()
	}
	return randSlice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if data == nil {
		return 0
	}
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := data[0]
	for i := range len(data) {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// // maxChunks returns the maximum number of elements in a chunks.
// func maxChunks(data []int) int {
// 	// ваш код здесь
// }

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	randSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// // ваш код здесь
	start := time.Now()
	max := maximum(randSlice)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	// fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// // ваш код здесь

	// fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
