package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateSort(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"Name", false},
		{"rating", true},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			ans := ValidateSort(tt.s)
			assert.Equal(t, ans, tt.want)
		})
	}
}
