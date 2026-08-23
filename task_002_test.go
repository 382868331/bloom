package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom002MinimumCapacity(t *testing.T){got:=New(1,1).Cap();if got!=uint(1){t.Fatalf("got %v want %v",got,uint(1))}}
