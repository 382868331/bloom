package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom014AddHashLoop(t *testing.T){got:=func() bool {f:=New(128,1);f.Add([]byte("x"));return f.Test([]byte("x"))}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom014AddHashLoopAdjacent(t *testing.T) {
	got := func() uint {f:=New(128,2);f.Add([]byte("y"));return f.BitSet().Count()}()
	want := uint(2)
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
