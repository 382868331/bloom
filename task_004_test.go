package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom004FromWordCapacity(t *testing.T){got:=From([]uint64{0,0},2).Cap();if got!=uint(128){t.Fatalf("got %v want %v",got,uint(128))}}
func TestTaskBloom004FromWordCapacityAdjacent(t *testing.T) {
	got := From([]uint64{0},2).Cap()
	want := uint(64)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
