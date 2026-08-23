package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom007LocationBaseLane(t *testing.T){got:=location([4]uint64{3,5,7,11},0);if got!=uint64(3){t.Fatalf("got %v want %v",got,uint64(3))}}
