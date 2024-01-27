package lesson1_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct { // добавляем слайс тестов
		name  string
		value float64
		want  float64
	}{
		{
			name:  "negative float",
			value: -3,
			want:  3,
		},
		{
			name:  "positive float",
			value: 3,
			want:  3,
		},
		{
			name:  "negative fractional float",
			value: -2.000001,
			want:  2.000001,
		},
		{
			name:  "negative close to zero float",
			value: -0.000000003,
			want:  0.000000003,
		},
		{
			name:  "positive float #2",
			value: 65.00013,
			want:  65.00013,
		},
	}
	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.value))
		})
	}
}
