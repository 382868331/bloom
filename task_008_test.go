package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom008FilterLocationModulo(t *testing.T){got:=New(10,2).location([4]uint64{19,3,5,7},0);if got!=uint(9){t.Fatalf("got %v want %v",got,uint(9))}}
