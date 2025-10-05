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
