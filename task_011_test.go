package bloom
import("math";"testing")
var _=math.Ceil
func TestTaskBloom011EstimateConstructorOrder(t *testing.T){got:=func() bool {m,k:=EstimateParameters(100,.01);f:=NewWithEstimates(100,.01);return f.Cap()==m&&f.K()==k}();if got!=true{t.Fatalf("got %v want %v",got,true)}}
