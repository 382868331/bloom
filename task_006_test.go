package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom006BaseHashOrder(t *testing.T){got:=func() bool { var d digest128;a,b,_,_:=d.sum256([]byte("x"));h:=baseHashes([]byte("x"));return h[0]==a&&h[1]==b }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom006BaseHashOrderAdjacent(t *testing.T) {
	got := func() bool { var d digest128;a,b,_,_:=d.sum256([]byte("y"));h:=baseHashes([]byte("y"));return h[0]==a&&h[1]==b }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
