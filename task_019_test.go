package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom019LocationTestPolarity(t *testing.T){got:=func() bool {f:=New(64,1);f.b.Set(3);return f.TestLocations([]uint64{3})}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
