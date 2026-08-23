package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom004FromWordCapacity(t *testing.T){got:=From([]uint64{0,0},2).Cap();if got!=uint(128){t.Fatalf("got %v want %v",got,uint(128))}}
