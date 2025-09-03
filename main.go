package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"time"
)

const (
	SIZE   = 100_000_000 // Размер массива для генерации случайных чисел
	CHUNKS = 8           // Количество чанков для параллельной обработки
)

// Инициализируем глобальный генератор случайных чисел с текущим временем
var src = rand.New(rand.NewSource(time.Now().UnixNano()))

// generateRandomElements генерирует массив случайных чисел заданного размера
func generateRandomElements(size int) ([]int, error) {

	// возвращаем пустой слайс и ошибку, если size <= 0
	if size <= 0 {
		return nil, fmt.Errorf("invalid slice size: %d. The size must be a positive integer greater than zero", size)
	}

	// Создаем слайс заданного размера
	randomElements := make([]int, size)

	// Заполняем слайс случайными числами
	for i := range randomElements {
		randomElements[i] = src.Int()
	}

	return randomElements, nil
}

// maximum находит максимальное значение в массиве
func maximum(data []int) (int, error) {

	/// Проверяем, что слайс не пустой
	if len(data) == 0 {
		return 0, fmt.Errorf("empty slice provided. The input slice must contain at least one element to determine the maximum value")
	}

	// Для слайса из одного элемента возвращаем его значение
	if len(data) == 1 {
		return data[0], nil
	}

	// Используем встроенную функцию для поиска максимума
	max := slices.Max(data)

	return max, nil
}

// maxChunks находит максимальное значение в массиве, используя параллельную обработку
func maxChunks(data []int) (int, error) {

	// Проверяем, что слайс не пустой
	if len(data) == 0 {
		return 0, fmt.Errorf("input slice cannot be empty. Please provide a non-empty slice of integers")
	}

	// Для слайса из одного элемента возвращаем его значение
	if len(data) == 1 {
		return data[0], nil
	}

	// Проверка, что длина слайса достаточна для разделения на чанки
	if len(data) < CHUNKS {
		return 0, fmt.Errorf("input slice length is insufficient to create required number of chunks")
	}

	// Создаем канал для хранения максимальных значений из каждого чанка и закрываем канал после использования
	maxValues := make(chan int, CHUNKS)
	defer close(maxValues)

	// Вычисляем размер каждого чанка
	step := len(data) / CHUNKS
	remainder := len(data) % CHUNKS

	// Разбиваем массив на куски и обрабатываем их параллельно
	for i := 0; i < CHUNKS; i++ {
		start := i * step
		end := start + step

		// Корректируем последний чанк, если есть остаток
		if i == CHUNKS-1 {
			end += remainder
		}

		// Запускаем горутину для обработки каждого чанка
		go func(start, end int) {
			maxValues <- slices.Max(data[start:end])
		}(start, end)
	}

	// Находим максимальное значение среди всех чанков
	var max int
	for i := 0; i < CHUNKS; i++ {
		val := <-maxValues
		if i == 0 || val > max {
			max = val
		}
	}

	return max, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)

	// Генерируем случайные числа
	randomElements, err := generateRandomElements(SIZE)
	if err != nil {
		log.Fatalf("failed to generate random numbers: %v", err)
	}

	// Измеряем время выполнения последовательного поиска максимума
	start := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	max, err := maximum(randomElements)
	if err != nil {
		log.Printf("error occurred while finding maximum value in single-threaded mode: %v", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n\n", max, elapsed.Microseconds())

	// Измеряем время выполнения параллельного поиска максимума
	start = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	max, err = maxChunks(randomElements)
	if err != nil {
		log.Printf("error occurred while finding maximum value in parallel mode: %v", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
