package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom009EstimatedBitsRound(t *testing.T){got:=func() bool {m,_:=EstimateParameters(100,.01);want:=uint(math.Ceil(-100*math.Log(.01)/math.Pow(math.Log(2),2)));return m==want}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
