package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom010EstimatedHashesRound(t *testing.T){got:=func() bool {m,k:=EstimateParameters(100,.01);return k==uint(math.Ceil(math.Log(2)*float64(m)/100))}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
