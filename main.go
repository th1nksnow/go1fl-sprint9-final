package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range size {
		data[i] = rnd.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
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

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	dataLength := len(data)
	chunkSize := dataLength / CHUNKS

	chunkMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		left := i * chunkSize
		right := left + chunkSize
		if i == CHUNKS-1 {
			right = dataLength
		}
		chunk := data[left:right]

		go func(chunkIndex int, chunk []int) {
			defer wg.Done()
			chunkMaxes[chunkIndex] = maximum(chunk)
		}(i, chunk)
	}

	wg.Wait()
	return maximum(chunkMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
