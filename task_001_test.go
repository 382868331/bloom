package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom001MaxSelection(t *testing.T){got:=max(1,5);if got!=uint(5){t.Fatalf("got %v want %v",got,uint(5))}}
