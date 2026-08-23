package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom014AddHashLoop(t *testing.T){got:=func() bool {f:=New(128,1);f.Add([]byte("x"));return f.Test([]byte("x"))}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
