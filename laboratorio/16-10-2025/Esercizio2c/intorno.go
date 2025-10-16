package main

import (
        "fmt"
        "math"
)

func main() {
  const EPSILON = 1e-5;

  var x, y float64;

  fmt.Scan(&x, &y);

  distanza := math.Sqrt(x*x + y*y);

  if distanza < EPSILON {
    fmt.Print("molto vicino all'origine");
  } else {
    fmt.Print("non vicino all'origine");
  }

}
