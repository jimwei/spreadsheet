package core
import "testing"
func TestStringUtil(t *testing.T)  {
    s := stringUtil.Number2String(123)
    t.Log(s)
}
func  BenchmarkUtil(t *testing.B)  {
    t.ResetTimer()
    for i:=0;i<10000;i++{
      stringUtil.Number2String(55657687999908986)  
    }
      
     t.StopTimer()
}