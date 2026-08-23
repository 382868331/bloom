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

func TestBloomProbabilityRangeOrder(t *testing.T) {
	if got := BloomProbabilityRangeOrder(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomProbabilityRangeOrderRegression(t *testing.T) {
	TestBloomProbabilityRangeOrder(t)
	TestBloomProbabilityRangeOrder(t)
}

func TestBloomCapacityProductOverflow(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := BloomCapacityProductOverflow(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestBloomCapacityProductOverflowRegression(t *testing.T) {
	TestBloomCapacityProductOverflow(t)
	TestBloomCapacityProductOverflow(t)
}

func TestBloomEscapedSeedSequence(t *testing.T) {
	got := BloomEscapedSeedSequence("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestBloomEscapedSeedSequenceRegression(t *testing.T) {
	TestBloomEscapedSeedSequence(t)
	TestBloomEscapedSeedSequence(t)
}

func TestBloomStableHashDedup(t *testing.T) {
	got := BloomStableHashDedup([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestBloomStableHashDedupRegression(t *testing.T) {
	TestBloomStableHashDedup(t)
	TestBloomStableHashDedup(t)
}

func TestBloomZeroHashCount(t *testing.T) {
	if got := BloomZeroHashCount([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestBloomZeroHashCountRegression(t *testing.T) {
	TestBloomZeroHashCount(t)
	TestBloomZeroHashCount(t)
}

func TestBloomUnicodeKeyPreview(t *testing.T) {
	got := BloomUnicodeKeyPreview("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestBloomUnicodeKeyPreviewRegression(t *testing.T) {
	TestBloomUnicodeKeyPreview(t)
	TestBloomUnicodeKeyPreview(t)
}

func TestBloomSerializedFlagSpace(t *testing.T) {
	got, err := BloomSerializedFlagSpace(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestBloomSerializedFlagSpaceRegression(t *testing.T) {
	TestBloomSerializedFlagSpace(t)
	TestBloomSerializedFlagSpace(t)
}

func TestBloomHashOffsetOverflow(t *testing.T) {
	if got := BloomHashOffsetOverflow(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomHashOffsetOverflowRegression(t *testing.T) {
	TestBloomHashOffsetOverflow(t)
	TestBloomHashOffsetOverflow(t)
}

func TestBloomFalsePositiveEndpoint(t *testing.T) {
	if got := BloomFalsePositiveEndpoint([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomFalsePositiveEndpointRegression(t *testing.T) {
	TestBloomFalsePositiveEndpoint(t)
	TestBloomFalsePositiveEndpoint(t)
}

func TestBloomBitsetCloneIsolation(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := BloomBitsetCloneIsolation(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestBloomBitsetCloneIsolationRegression(t *testing.T) {
	TestBloomBitsetCloneIsolation(t)
	TestBloomBitsetCloneIsolation(t)
}

func TestBloomReverseKeyUnicode(t *testing.T) {
	if got := BloomReverseKeyUnicode("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestBloomReverseKeyUnicodeRegression(t *testing.T) {
	TestBloomReverseKeyUnicode(t)
	TestBloomReverseKeyUnicode(t)
}

func TestBloomPartitionTailBits(t *testing.T) {
	got := BloomPartitionTailBits([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestBloomPartitionTailBitsRegression(t *testing.T) {
	TestBloomPartitionTailBits(t)
	TestBloomPartitionTailBits(t)
}

func TestBloomEmptyUnionList(t *testing.T) {
	if got := BloomEmptyUnionList(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestBloomEmptyUnionListRegression(t *testing.T) {
	TestBloomEmptyUnionList(t)
	TestBloomEmptyUnionList(t)
}

func TestBloomConcurrentBitSet(t *testing.T) {
	if got := BloomConcurrentBitSet(64); got != 64 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomConcurrentBitSetRegression(t *testing.T) {
	TestBloomConcurrentBitSet(t)
	TestBloomConcurrentBitSet(t)
}

func TestBloomCanceledBulkAdd(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := BloomCanceledBulkAdd(ctx, 20); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestBloomCanceledBulkAddRegression(t *testing.T) {
	TestBloomCanceledBulkAdd(t)
	TestBloomCanceledBulkAdd(t)
}

func TestBloomUnionSizeError(t *testing.T) {
	base := errors.New("root")
	if got := BloomUnionSizeError(base); !errors.Is(got, base) {
		t.Fatalf("chain lost: %v", got)
	}
}

func TestBloomUnionSizeErrorRegression(t *testing.T) {
	TestBloomUnionSizeError(t)
	TestBloomUnionSizeError(t)
}

func TestBloomDecoderBufferCleanup(t *testing.T) {
	active = 0
	BloomDecoderBufferCleanup(true)
	if active != 0 {
		t.Fatalf("active=%d", active)
	}
}

func TestBloomDecoderBufferCleanupRegression(t *testing.T) {
	TestBloomDecoderBufferCleanup(t)
	TestBloomDecoderBufferCleanup(t)
}

func TestBloomSparseBitIterator(t *testing.T) {
	got := BloomSparseBitIterator([]int{2, 4, 5, 6})
	if !reflect.DeepEqual(got, []int{5}) {
		t.Fatalf("got %v", got)
	}
}

func TestBloomSparseBitIteratorRegression(t *testing.T) {
	TestBloomSparseBitIterator(t)
	TestBloomSparseBitIterator(t)
}

func TestBloomTestAddReturnOrder(t *testing.T) {
	v, ok := BloomTestAddReturnOrder(nil)
	if ok || v != 0 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestBloomTestAddReturnOrderRegression(t *testing.T) {
	TestBloomTestAddReturnOrder(t)
	TestBloomTestAddReturnOrder(t)
}
