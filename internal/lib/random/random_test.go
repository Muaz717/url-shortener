package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T){
	tests := []struct{
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 5",
			size: 5,
		},
		{
			name: "size = 10",
			size: 10,
		},
		{
			name: "size = 20",
			size: 20,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			str1 := NewRandomString(test.size)
			str2 := NewRandomString(test.size)

			assert.Len(t, str1, test.size)
			assert.Len(t, str2, test.size)
			
			assert.NotEqual(t, str1, str2)
		})
	}
}
