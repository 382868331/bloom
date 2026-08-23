package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom008FilterLocationModulo(t *testing.T){got:=New(10,2).location([4]uint64{19,3,5,7},0);if got!=uint(9){t.Fatalf("got %v want %v",got,uint(9))}}
func TestTaskBloom008FilterLocationModuloAdjacent(t *testing.T) {
	got := New(8,2).location([4]uint64{15,3,5,7},0)
	want := uint(7)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
