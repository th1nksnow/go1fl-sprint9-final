package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
		wantNil  bool
	}{
		{
			name:     "Positive size",
			size:     5,
			expected: 5,
			wantNil:  false,
		},
		{
			name:     "Zero size",
			size:     0,
			expected: 0,
			wantNil:  false,
		},
		{
			name:     "Niggative size",
			size:     -5,
			expected: 0,
			wantNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if tt.wantNil {
				assert.Nil(t, result)
				return
			}

			assert.Equal(t, tt.expected, len(result))
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Positive numbers",
			input: []int{1, 5, 3, 9, 2},
			want:  9,
		},
		{
			name:  "Mixed positive and negative numbers",
			input: []int{10, 0, 5, 3, 8},
			want:  10,
		},
		{
			name:  "Single element",
			input: []int{42},
			want:  42,
		},
		{
			name:  "All equal numbers",
			input: []int{7, 7, 7, 7, 7},
			want:  7,
		},
		{
			name:  "Empty slice",
			input: []int{},
			want:  0,
		},
		{
			name:  "Nil slice",
			input: nil,
			want:  0,
		},
		{
			name:  "With zero values",
			input: []int{0, 1, 0, 5, 0},
			want:  5,
		},
		{
			name:  "Large numbers",
			input: []int{1000000, 999999, 1000001},
			want:  1000001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Single element",
			input: []int{42},
			want:  42,
		},
		{
			name:  "Empty slice",
			input: []int{},
			want:  0,
		},
		{
			name:  "Nil slice",
			input: nil,
			want:  0,
		},
		{
			name:  "With <= CHUNKS elements",
			input: []int{0, 1, 0, 5, 0},
			want:  5,
		},
		{
			name:  "Regular length",
			input: []int{0, 1, 0, 5, 0, 14, 228, 11, 5, 100500},
			want:  100500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)

			assert.Equal(t, tt.want, result)
		})
	}
}
