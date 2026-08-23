package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom017CopyContent(t *testing.T){got:=func() bool {f:=New(128,2);f.AddString("x");return f.Copy().TestString("x")}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom017CopyContentAdjacent(t *testing.T) {
	got := func() bool {f:=New(128,2);f.AddString("y");c:=f.Copy();return c.Cap()==f.Cap()&&c.TestString("y")}()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
