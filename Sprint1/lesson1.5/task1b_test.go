package lesson1_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{
			name: "Simple user",
			user: User{"Sergey", "Voronov"},
			want: "Sergey Voronov",
		},
		{
			name: "Foreign user",
			user: User{"Mike", "Johnson"},
			want: "Mike Johnson",
		},
		{
			name: "Disney user",
			user: User{"Scrooge", "McDuck"},
			want: "Scrooge McDuck",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, test.user.FullName())
		})
	}
}
