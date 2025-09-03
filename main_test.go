package main

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест проверяет корректность работы функции generateRandomElements при различных входных параметрах
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		size    int    // Передаваемый размер
		wantErr bool   // Ожидаем ли ошибку
		wantLen int    // Ожидаемая длина слайса
	}{
		// Тест на генерацию пустого массива (размер 0)
		{
			name:    "ZeroSizeInput",
			size:    0,
			wantErr: true,
			wantLen: 0,
		},
		// Тест на отрицательный размер (невалидный ввод)
		{
			name:    "NegativeSizeInput",
			size:    -5,
			wantErr: true,
			wantLen: 0,
		},
		// Тест на положительный размер (оптимальный сценарий)
		{
			name:    "PositiveSizeInput",
			size:    10,
			wantErr: false,
			wantLen: 10,
		},
		// Тест на генерацию одного элемента
		{
			name:    "SingleElementInput",
			size:    1,
			wantErr: false,
			wantLen: 1,
		},
		// Тест на генерацию большого массива
		{
			name:    "LargeSizeInput",
			size:    1000000000,
			wantErr: false,
			wantLen: 1000000000,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			randomElements, err := generateRandomElements(tt.size)

			// Проверка наличия/отсутствия ошибки
			if tt.wantErr {
				require.Error(t, err, "expected error for invalid input size but got none")
				return
			}
			require.NoError(t, err, "unexpected error occurred during element generation")

			// Проверка корректности результата
			assert.Equal(t, tt.wantLen, len(randomElements), "generated slice length does not match expected value")

			// Измерение времени выполнения теста
			t.Logf("Тест %q выполнен за %v", tt.name, time.Since(start))
		})
	}
}

// TestMaximum - тестовая функция для проверки корректности работы функции поиска максимального значения в слайсе
func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		data    []int  // Передаваемый слайс для обработки
		wantErr bool   // Ожидаем ли ошибку при выполнении
		wantInt int    // Ожидаемое максимальное значение
	}{
		// Тест на пустой входной массив
		{
			name:    "EmptyInput",
			data:    []int{},
			wantErr: true,
			wantInt: 0,
		},
		// Тест на массив с одним элементом
		{
			name:    "SingleElement",
			data:    []int{1},
			wantErr: false,
			wantInt: 1,
		},
		// Тест на большой массив положительных чисел
		{
			name:    "LargeInput",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantErr: false,
			wantInt: 10,
		},
		// Тест на массив отрицательных чисел
		{
			name:    "NegativeNumbers",
			data:    []int{-1, -2, -3, -4},
			wantErr: false,
			wantInt: -1,
		},
		// Тест на массив с разными типами чисел
		{
			name:    "MixedNumbers",
			data:    []int{-10, 0, 5, -3, 8},
			wantErr: false,
			wantInt: 8,
		},
		// Тест на массив с одинаковыми элементами
		{
			name:    "AllIdentical",
			data:    []int{5, 5, 5, 5},
			wantErr: false,
			wantInt: 5,
		},
		// Тест на максимальное значение int
		{
			name:    "MaxIntValue",
			data:    []int{math.MaxInt32, 1, 2},
			wantErr: false,
			wantInt: math.MaxInt32,
		},
		// Тест на минимальное значение int
		{
			name:    "MinIntValue",
			data:    []int{math.MinInt32, -1, -2},
			wantErr: false,
			wantInt: -1,
		},
		// Тест на массив из нулей
		{
			name:    "AllZeros",
			data:    []int{0, 0, 0, 0},
			wantErr: false,
			wantInt: 0,
		},
		// Тест на массив с повторяющимися максимальными значениями
		{
			name:    "DuplicateMaxValues",
			data:    []int{3, 7, 7, 2, 7},
			wantErr: false,
			wantInt: 7,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			max, err := maximum(tt.data)

			// Проверка наличия/отсутствия ошибки
			if tt.wantErr {
				require.Error(t, err, "Ожидалась ошибка, но её нет")
				return
			}
			require.NoError(t, err, "Не ожидалось получение ошибки")

			// Проверка корректности результата
			assert.Equal(t, tt.wantInt, max, "Результат не соответствует ожидаемому")

			// Измерение времени выполнения теста
			t.Logf("Тест %q выполнен за %v", tt.name, time.Since(start))
		})
	}
}

// TestMaxChunks - тестовая функция для проверки корректности работы функции поиска максимального значения с разбиением на чанки (с использованием горутин)
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name    string // Наименование теста
		data    []int  // Передаваемый слайс для обработки
		wantErr bool   // Ожидаем ли ошибку при выполнении
		wantInt int    // Ожидаемое максимальное значение
	}{
		// Базовые тесты
		{
			name:    "Размер 0",
			data:    []int{},
			wantErr: true,
			wantInt: 0,
		},
		{
			name:    "Размер 1",
			data:    []int{1},
			wantErr: false,
			wantInt: 1,
		},
		{
			name:    "Длина равна CHUNKS",
			data:    make([]int, 8),
			wantErr: false,
			wantInt: 0,
		},
		{
			name:    "Длина меньше CHUNKS",
			data:    []int{1, 2, 3, 4},
			wantErr: true,
			wantInt: 0,
		},

		// Тесты с разными значениями CHUNKS
		{
			name:    "CHUNKS = 1",
			data:    []int{1, 2, 3, 4, 5},
			wantErr: true,
			wantInt: 5,
		},
		{
			name:    "CHUNKS = 2",
			data:    []int{1, 2, 3, 4, 5, 6},
			wantErr: true,
			wantInt: 6,
		},

		// Граничные случаи
		{
			name:    "Максимум в начале",
			data:    []int{100, 1, 2, 3, 4, 5, 6, 7},
			wantErr: false,
			wantInt: 100,
		},
		{
			name:    "Максимум в конце",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 100},
			wantErr: false,
			wantInt: 100,
		},

		// Сложные случаи
		{
			name:    "Неравномерное распределение",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			wantErr: false,
			wantInt: 13,
		},
		{
			name:    "Большой массив",
			data:    make([]int, 10000),
			wantErr: false,
			wantInt: 0,
		},

		// Тесты с разными типами данных
		{
			name:    "Максимальное int",
			data:    []int{math.MaxInt32, 1, 2, 8, 15, 84, 99, 91},
			wantErr: false,
			wantInt: math.MaxInt32,
		},
		{
			name:    "Минимальное int",
			data:    []int{math.MinInt32, -1, -2, -8, -15, -84, -99, -91},
			wantErr: false,
			wantInt: -1,
		},
		{
			name:    "С нулями",
			data:    []int{0, 0, 0, 0, 0, 0, 0, 0},
			wantErr: false,
			wantInt: 0,
		},
		{
			name:    "Повторяющиеся максимумы",
			data:    []int{3, 7, 7, 2, 7, 5, 4, 3},
			wantErr: false,
			wantInt: 7,
		},
		{
			name:    "Чередование максимумов",
			data:    []int{10, 1, 10, 1, 10, 1, 10, 1},
			wantErr: false,
			wantInt: 10,
		},
		{
			name:    "Все отрицательные",
			data:    []int{-10, -20, -30, -40, -50, -60, -70, -80},
			wantErr: false,
			wantInt: -10,
		},
		{
			name:    "Переменный шаг",
			data:    []int{1, 10, 2, 20, 3, 30, 4, 40},
			wantErr: false,
			wantInt: 40,
		},
		{
			name:    "Экспоненциальный рост",
			data:    []int{1, 2, 4, 8, 16, 32, 64, 128},
			wantErr: false,
			wantInt: 128,
		},
		{
			name:    "Линейный рост",
			data:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			wantErr: false,
			wantInt: 8,
		},
		{
			name:    "Линейное убывание",
			data:    []int{8, 7, 6, 5, 4, 3, 2, 1},
			wantErr: false,
			wantInt: 8,
		},
	}

	// Проход по всем тестовым случаям
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			// Вызов тестируемой функции
			max, err := maxChunks(tt.data)

			// Проверка наличия/отсутствия ошибки
			if tt.wantErr {
				require.Error(t, err, "Ожидалась ошибка, но её нет")
				return
			}
			require.NoError(t, err, "Не ожидалось получение ошибки")

			// Проверка корректности результата
			assert.Equal(t, tt.wantInt, max, "Тест %q: результат не соответствует ожидаемому. Ожидалось: %d, получено: %d", tt.name, tt.wantInt, max)

			// Измерение времени выполнения теста
			t.Logf("Тест %q выполнен за %v", tt.name, time.Since(start))
		})
	}
}
