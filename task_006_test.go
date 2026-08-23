package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom006BaseHashOrder(t *testing.T){got:=func() bool { var d digest128;a,b,_,_:=d.sum256([]byte("x"));h:=baseHashes([]byte("x"));return h[0]==a&&h[1]==b }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
