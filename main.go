package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000 // Размер слайса для генерации случайных чисел
	CHUNKS = 8           // Количество чанков для параллельной обработки
)

// Инициализируем глобальный генератор случайных чисел с текущим временем
var src = rand.New(rand.NewSource(time.Now().UnixNano()))

// generateRandomElements генерирует слайс случайных чисел заданного размера
func generateRandomElements(size int) []int {

	// Логируем ошибку и возвращаем nil при недопустимом размере
	if size <= 0 {
		log.Printf("invalid slice size: %d. The size must be a positive integer greater than zero", size)
		return nil
	}

	// Создаем слайс заданного размера
	randomElements := make([]int, size)

	// Заполняем слайс случайными числами
	for i := range randomElements {
		randomElements[i] = src.Int()
	}

	return randomElements
}

// maximum находит максимальное значение в слайсе
func maximum(data []int) int {

	// Проверяем, что входной слайс не пустой. Если слайс пуст, логируем ошибку и возвращаем 0
	if len(data) == 0 {
		log.Printf("empty slice provided. The input slice must contain at least one element to determine the maximum value")
		return 0
	}

	// Инициализируем переменную max первым элементом слайса
	max := data[0]

	// Проходим по всем элементам слайса и сравниваем каждый элемент с текущим максимальным значением
	for _, num := range data {
		if num > max {
			max = num
		}
	}

	return max
}

// maxChunks находит максимальное значение в слайсе, используя параллельную обработку
func maxChunks(data []int) int {

	// Проверяем, что входной слайс не пустой. Если слайс пуст, логируем ошибку и возвращаем 0
	if len(data) == 0 {
		log.Printf("empty slice provided. The input slice must contain at least one element to determine the maximum value")
		return 0
	}

	// Проверяем, что длина слайса достаточна для разделения на чанки. Если данных недостаточно, вызываем функцию maximum
	if len(data) < CHUNKS {
		maximum(data)
	}

	var wg sync.WaitGroup

	// Создаем слайс для хранения максимальных значений из каждого чанка
	maxValues := make([]int, CHUNKS)

	// Вычисляем размер каждого чанка
	step := len(data) / CHUNKS
	remainder := len(data) % CHUNKS

	// Разбиваем слайс на куски и обрабатываем их параллельно
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * step
		end := start + step

		// Корректируем последний чанк, если есть остаток
		if i == CHUNKS-1 {
			end += remainder
		}

		// Запускаем горутину для обработки каждого чанка
		go func(start, end int) {
			defer wg.Done()
			maxValues[i] = maximum(data[start:end])
		}(start, end)
	}

	// Ждем завершения всех горутин перед возвратом результата
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)

	// Генерируем случайные числа
	randomElements := generateRandomElements(SIZE)

	// Измеряем время выполнения последовательного поиска максимума
	start := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	max := maximum(randomElements)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n\n", max, elapsed.Microseconds())

	// Измеряем время выполнения параллельного поиска максимума
	start = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	max = maxChunks(randomElements)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
