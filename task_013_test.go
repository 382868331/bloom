package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom013HashAccessor(t *testing.T){got:=New(64,3).K();if got!=uint(3){t.Fatalf("got %v want %v",got,uint(3))}}
