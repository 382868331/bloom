package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom005FromHashCount(t *testing.T){got:=FromWithM([]uint64{0},64,3).K();if got!=uint(3){t.Fatalf("got %v want %v",got,uint(3))}}
func TestTaskBloom005FromHashCountAdjacent(t *testing.T) {
	got := FromWithM([]uint64{0},64,1).K()
	want := uint(1)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
