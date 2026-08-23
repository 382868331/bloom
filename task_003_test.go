package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom003MinimumHashCount(t *testing.T){got:=New(8,1).K();if got!=uint(1){t.Fatalf("got %v want %v",got,uint(1))}}
func TestTaskBloom003MinimumHashCountAdjacent(t *testing.T) {
	got := New(8,0).K()
	want := uint(1)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
