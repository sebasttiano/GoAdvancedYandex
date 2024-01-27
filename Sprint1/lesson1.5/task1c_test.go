package lesson1_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFamily_AddNew(t *testing.T) {
	family := Family{}
	tests := []struct {
		name         string
		relationship Relationship
		person       Person
		want         error
	}{
		{
			name:         "First father",
			relationship: Father,
			person:       Person{FirstName: "Mike", LastName: "Johnson", Age: 40},
			want:         nil,
		},
		{
			name:         "Child",
			relationship: Child,
			person:       Person{FirstName: "Ann", LastName: "Johnson", Age: 15},
			want:         nil,
		},
		{
			name:         "Second father",
			relationship: Father,
			person:       Person{FirstName: "Bob", LastName: "Johnson", Age: 39},
			want:         ErrRelationshipAlreadyExists,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, family.AddNew(test.relationship, test.person))
		})
	}
}
