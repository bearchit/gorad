package gorad_test

import (
	"github.com/bearchit/gorad"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b, want int
	}{
		{1, 2, 1},
		{2, 1, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := gorad.Min(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{2, 1, 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := gorad.Max(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
