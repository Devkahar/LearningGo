package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	addCases := []struct {
		name          string
		expected      int
		valA          int
		valB          int
		expectedError error
	}{
		{
			name:          "Testing correct addition of 2 and 5",
			expected:      7,
			valA:          2,
			valB:          5,
			expectedError: nil,
		},
		{
			name:          "Testing addition of 2 and -5",
			expected:      0,
			valA:          2,
			valB:          -5,
			expectedError: errors.New("Cannot Add Negative value."),
		},
	}

	for _, test := range addCases {
		t.Run(test.name, func(t *testing.T) {
			gotVal, gotError := Add(test.valA, test.valB)
			assert.Equal(t, test.expected, gotVal)
			assert.Equal(t, test.expectedError, gotError)
		})
	}
}
