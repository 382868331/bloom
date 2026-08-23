package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom005FromHashCount(t *testing.T){got:=FromWithM([]uint64{0},64,3).K();if got!=uint(3){t.Fatalf("got %v want %v",got,uint(3))}}
