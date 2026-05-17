package logic

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{
			name:        "Положительные числа",
			a:           10.0,
			b:           2.0,
			expected:    5.0,
			expectError: false,
		},
		{
			name:        "Деление на ноль",
			a:           5.0,
			b:           0.0,
			expected:    0,
			expectError: true,
		},
		{
			name:        "Отрицательные числа",
			a:           -10.0,
			b:           2.0,
			expected:    -5.0,
			expectError: false,
		},
		{
			name:        "Отрицательный делитель",
			a:           10.0,
			b:           -2.0,
			expected:    -5.0,
			expectError: false,
		},
		{
			name:        "Два отрицательных",
			a:           -10.0,
			b:           -2.0,
			expected:    5.0,
			expectError: false,
		},
		{
			name:        "Дробные числа",
			a:           7.5,
			b:           2.5,
			expected:    3.0,
			expectError: false,
		},
		{
			name:        "Деление на ноль с отрицательным числом",
			a:           -5.0,
			b:           0.0,
			expected:    0,
			expectError: true,
		},
		{
			name:        "Ноль делить на число",
			a:           0.0,
			b:           5.0,
			expected:    0.0,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError && err == nil {
				t.Errorf("Divide(%f, %f) ожидалась ошибка, но её нет", tt.a, tt.b)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Divide(%f, %f) не ожидалась ошибка, но получена: %v", tt.a, tt.b, err)
			}

			if !tt.expectError && result != tt.expected {
				t.Errorf("Divide(%f, %f) = %f, ожидалось %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
} 