package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "обычный размер", size: 100},
		{name: "один элемент", size: 1},
		{name: "нулевой размер", size: 0},
		{name: "отрицательный размер", size: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			expectedSize := tt.size
			if expectedSize < 0 {
				expectedSize = 0
			}
			if len(data) != expectedSize {
				t.Fatalf("получена длина %d, ожидалась %d", len(data), expectedSize)
			}
			for _, value := range data {
				if value <= 0 {
					t.Fatalf("получено неположительное число %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "обычный слайс", data: []int{3, 8, 2, 10, 4}, want: 10},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "nil", data: nil, want: 0},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "одинаковые элементы", data: []int{5, 5, 5}, want: 5},
		{name: "отрицательные элементы", data: []int{-8, -2, -10}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Fatalf("получено %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "делится на восемь", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "есть остаток", data: []int{1, 2, 30, 4, 5, 6, 7, 8, 9, 10}, want: 30},
		{name: "меньше восьми элементов", data: []int{3, 12, 5}, want: 12},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "nil", data: nil, want: 0},
		{name: "отрицательные элементы", data: []int{-8, -2, -10}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Fatalf("получено %d, ожидалось %d", got, tt.want)
			}
		})
	}
}
