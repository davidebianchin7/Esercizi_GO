package main
import "fmt"

func main(){
  var n int
  fmt.Scan(&n)

  for i := 1; ;i++{
    x := i * i

    if x > n{
      break
    }

    fmt.Println(x)
  }
}
