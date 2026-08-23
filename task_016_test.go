package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom016MergeHashGuard(t *testing.T){got:=func() bool {f:=New(64,2);g:=New(64,2);return f.Merge(g)==nil}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom016MergeHashGuardAdjacent(t *testing.T) {
	got := func() bool {f:=New(128,3);g:=New(128,3);return f.Merge(g)==nil}()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
