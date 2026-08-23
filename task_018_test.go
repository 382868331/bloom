package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom018MembershipPolarity(t *testing.T){got:=func() bool {f:=New(128,2);f.AddString("x");return f.TestString("x")}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom018MembershipPolarityAdjacent(t *testing.T) {
	got := func() bool {f:=New(128,3);f.AddString("y");return f.Test([]byte("y"))}()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
