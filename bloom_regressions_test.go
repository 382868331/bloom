package bloom

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"unicode/utf8"
)

var (
	_ = context.Background
	_ = errors.Is
	_ = reflect.DeepEqual
	_ = utf8.ValidString
)

func TestBloomNormalizeBounds(t *testing.T) {
	if got := BloomNormalizeBounds(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomNormalizeBoundsRegression(t *testing.T) {
	TestBloomNormalizeBounds(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomNormalizeBounds(t)
}
