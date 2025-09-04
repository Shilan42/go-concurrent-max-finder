package main

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Тест проверяет корректность работы функции generateRandomElements при различных входных параметрах
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		size    int    // Передаваемый размер
		wantLen int    // Ожидаемая длина слайса
		wantNil bool   // Ожидаем ли nil результат
	}{
		// Тест на генерацию пустого слайса (размер 0)
		{
			name:    "ZeroSizeInput",
			size:    0,
			wantLen: 0,
			wantNil: true,
		},
		// Тест на отрицательный размер (невалидный ввод)
		{
			name:    "NegativeSizeInput",
			size:    -5,
			wantLen: 0,
			wantNil: true,
		},
		// Тест на положительный размер (оптимальный сценарий)
		{
			name:    "PositiveSizeInput",
			size:    10,
			wantLen: 10,
			wantNil: false,
		},
		// Тест на генерацию одного элемента
		{
			name:    "SingleElementInput",
			size:    1,
			wantLen: 1,
			wantNil: false,
		},
		// Тест на генерацию большого слайса
		{
			name:    "LargeSizeInput",
			size:    1000000000,
			wantLen: 1000000000,
			wantNil: false,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			randomElements := generateRandomElements(tt.size)

			// Проверка результата для случаев, где ожидается nil
			if tt.wantNil {
				assert.Nil(t, randomElements, "expected nil result but got non-nil value")
				return
			}

			// Проверка корректности результата для не-nil случаев
			assert.Equal(t, tt.wantLen, len(randomElements), "generated slice length does not match expected value")
			assert.NotNil(t, randomElements, "expected non-nil result but got nil value")

			// Измерение времени выполнения теста
			t.Logf("Test %q executed in %v", tt.name, time.Since(start))
		})
	}
}

// TestMaximum - тестовая функция для проверки корректности работы функции поиска максимального значения в слайсе
func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		data    []int  // Передаваемый слайс для обработки
		wantInt int    // Ожидаемое максимальное значение
	}{
		// Тест на пустой входной слайс
		{
			name:    "EmptyInput",
			data:    []int{},
			wantInt: 0,
		},
		// Тест на слайс с одним элементом
		{
			name:    "SingleElement",
			data:    []int{1},
			wantInt: 1,
		},
		// Тест на большой слайс положительных чисел
		{
			name:    "LargeInput",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantInt: 10,
		},
		// Тест на слайс с разными типами чисел
		{
			name:    "MixedNumbers",
			data:    []int{-10, 0, 5, -3, 8},
			wantInt: 8,
		},
		// Тест на слайс с одинаковыми элементами
		{
			name:    "AllIdentical",
			data:    []int{5, 5, 5, 5},
			wantInt: 5,
		},
		// Тест на максимальное значение int
		{
			name:    "MaxIntValue",
			data:    []int{math.MaxInt32, 1, 2},
			wantInt: math.MaxInt32,
		},
		// Тест на минимальное значение int
		{
			name:    "MinIntValue",
			data:    []int{math.MinInt32, -1, -2},
			wantInt: -1,
		},
		// Тест на слайс из нулей
		{
			name:    "AllZeros",
			data:    []int{0, 0, 0, 0},
			wantInt: 0,
		},
		// Тест на слайс с повторяющимися максимальными значениями
		{
			name:    "DuplicateMaxValues",
			data:    []int{3, 7, 7, 2, 7},
			wantInt: 7,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			max := maximum(tt.data)

			// Проверка корректности результата
			assert.Equal(t, tt.wantInt, max, "test %q: result does not match expected value. Expected: %d, got: %d", tt.name, tt.wantInt, max)

			// Измерение времени выполнения теста
			t.Logf("Test %q executed in %v", tt.name, time.Since(start))
		})
	}
}

// TestMaxChunks - тестовая функция для проверки корректности работы функции поиска максимального значения с разбиением на чанки (с использованием горутин)
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		data    []int  // Передаваемый слайс для обработки
		wantInt int    // Ожидаемое максимальное значение
	}{
		// Базовые тесты
		{
			name:    "ZeroSizeInput",
			data:    []int{},
			wantInt: 0,
		},
		{
			name:    "SingleElement",
			data:    []int{1},
			wantInt: 1,
		},
		{
			name:    "LengthEqualsChunkSize",
			data:    make([]int, 8),
			wantInt: 0,
		},
		{
			name:    "LengthLessThanChunkSize1",
			data:    []int{1, 2, 3, 4},
			wantInt: 4,
		},
		{
			name:    "LengthLessThanChunkSize2",
			data:    []int{1, 2, 3, 4, 5},
			wantInt: 5,
		},
		{
			name:    "LengthLessThanChunkSize3",
			data:    []int{1, 2, 3, 4, 5, 6},
			wantInt: 6,
		},

		// Граничные случаи
		{
			name:    "MaxValueAtStart",
			data:    []int{100, 1, 2, 3, 4, 5, 6, 7},
			wantInt: 100,
		},
		{
			name:    "MaxValueAtEnd",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 100},
			wantInt: 100,
		},

		// Сложные случаи
		{
			name:    "UnevenDistribution",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			wantInt: 13,
		},
		{
			name:    "LargeInput",
			data:    make([]int, 10000),
			wantInt: 0,
		},

		// Тесты с разными типами данных
		{
			name:    "MaxIntValue",
			data:    []int{math.MaxInt32, 1, 2, 8, 15, 84, 99, 91},
			wantInt: math.MaxInt32,
		},
		{
			name:    "MinIntValue",
			data:    []int{math.MinInt32, -1, -2, -8, -15, -84, -99, -91},
			wantInt: -1,
		},
		{
			name:    "AllZeros",
			data:    []int{0, 0, 0, 0, 0, 0, 0, 0},
			wantInt: 0,
		},
		{
			name:    "DuplicateMaxValues",
			data:    []int{3, 7, 7, 2, 7, 5, 4, 3},
			wantInt: 7,
		},
		{
			name:    "AlternatingMaxValues",
			data:    []int{10, 1, 10, 1, 10, 1, 10, 1},
			wantInt: 10,
		},
		{
			name:    "VariableStep",
			data:    []int{1, 10, 2, 20, 3, 30, 4, 40},
			wantInt: 40,
		},
		{
			name:    "ExponentialGrowth",
			data:    []int{1, 2, 4, 8, 16, 32, 64, 128},
			wantInt: 128,
		},
		{
			name:    "LinearGrowth",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			wantInt: 8,
		},
		{
			name:    "LinearDecay",
			data:    []int{8, 7, 6, 5, 4, 3, 2, 1},
			wantInt: 8,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			max := maxChunks(tt.data)

			// Проверка корректности результата
			assert.Equal(t, tt.wantInt, max, "test %q: result does not match expected value. Expected: %d, got: %d", tt.name, tt.wantInt, max)

			// Измерение времени выполнения теста
			t.Logf("Test %q executed in %v", tt.name, time.Since(start))
		})
	}
}
