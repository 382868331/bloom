package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom019LocationTestPolarity(t *testing.T){got:=func() bool {f:=New(64,1);f.b.Set(3);return f.TestLocations([]uint64{3})}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskBloom019LocationTestPolarityAdjacent(t *testing.T) {
	got := func() bool {f:=New(64,1);f.b.Set(5);f.b.Set(7);return f.TestLocations([]uint64{5,7})}()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
