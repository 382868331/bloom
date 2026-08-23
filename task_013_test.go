package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom013HashAccessor(t *testing.T){got:=New(64,3).K();if got!=uint(3){t.Fatalf("got %v want %v",got,uint(3))}}
func TestTaskBloom013HashAccessorAdjacent(t *testing.T) {
	got := New(128,5).K()
	want := uint(5)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
