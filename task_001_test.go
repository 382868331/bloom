package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom001MaxSelection(t *testing.T){got:=max(1,5);if got!=uint(5){t.Fatalf("got %v want %v",got,uint(5))}}
func TestTaskBloom001MaxSelectionAdjacent(t *testing.T) {
	got := max(7,2)
	want := uint(7)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
