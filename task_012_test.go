package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom012CapacityAccessor(t *testing.T){got:=New(64,3).Cap();if got!=uint(64){t.Fatalf("got %v want %v",got,uint(64))}}
