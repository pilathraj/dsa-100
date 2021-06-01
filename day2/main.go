package main
import "fmt"

func findNumber(a []int, c int) []int{
  var o []int
  n := len(a)
  for i:=0; i<n-(c-1); i++{
    if a[i] == a[i+1] && a[i]== a[i+2]{
      o = append(o, a[i])
    }
  }
  return o
}
func main(){
  a := []int{4, 5, 5, 5, 3, 8,22,22,22,5}
  fmt.Println(findNumber(a,3))
}
