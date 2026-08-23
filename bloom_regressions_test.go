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

func TestBloomPartitionValues(t *testing.T) {
	if got := BloomPartitionValues([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestBloomPartitionValuesRegression(t *testing.T) {
	TestBloomPartitionValues(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomPartitionValues(t)
}

func TestBloomTruncateLabel(t *testing.T) {
	got := BloomTruncateLabel("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestBloomTruncateLabelRegression(t *testing.T) {
	TestBloomTruncateLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomTruncateLabel(t)
}

func TestBloomParseBooleanOption(t *testing.T) {
	got, err := BloomParseBooleanOption(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestBloomParseBooleanOptionRegression(t *testing.T) {
	TestBloomParseBooleanOption(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomParseBooleanOption(t)
}

func TestBloomBoundedBackoff(t *testing.T) {
	if got := BloomBoundedBackoff(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomBoundedBackoffRegression(t *testing.T) {
	TestBloomBoundedBackoff(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomBoundedBackoff(t)
}

func TestBloomSelectUpperQuantile(t *testing.T) {
	if got := BloomSelectUpperQuantile([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomSelectUpperQuantileRegression(t *testing.T) {
	TestBloomSelectUpperQuantile(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomSelectUpperQuantile(t)
}

func TestBloomCloneNestedState(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := BloomCloneNestedState(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestBloomCloneNestedStateRegression(t *testing.T) {
	TestBloomCloneNestedState(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestBloomCloneNestedState(t)
}

func TestBloomReverseUnicodeLabel(t *testing.T) {
	if got := BloomReverseUnicodeLabel("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}
