package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom020ClearAllBits(t *testing.T){got:=func() bool {f:=New(64,2);f.AddString("x");f.ClearAll();return f.BitSet().Count()==0}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
