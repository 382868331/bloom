package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom003MinimumHashCount(t *testing.T){got:=New(8,1).K();if got!=uint(1){t.Fatalf("got %v want %v",got,uint(1))}}
