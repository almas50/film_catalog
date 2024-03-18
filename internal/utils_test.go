package internal

import (
	"fmt"
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
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := ValidateSort(tt.s)
			assert.Equal(t, ans, tt.want)
		})
	}
}
