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

func TestBloomSaturatingAdd(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := BloomSaturatingAdd(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestBloomSaturatingAddRegression(t *testing.T) {
	TestBloomSaturatingAdd(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomSaturatingAdd(t)
}

func TestBloomSplitEscapedTokens(t *testing.T) {
	got := BloomSplitEscapedTokens("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestBloomSplitEscapedTokensRegression(t *testing.T) {
	TestBloomSplitEscapedTokens(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomSplitEscapedTokens(t)
}

func TestBloomStableUnique(t *testing.T) {
	got := BloomStableUnique([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestBloomStableUniqueRegression(t *testing.T) {
	TestBloomStableUnique(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomStableUnique(t)
}
