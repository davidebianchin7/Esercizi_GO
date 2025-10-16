package main;

import (
        "fmt"
        "math"
)

func main() {
  const EPSILON = 1e-06;

  var x1, y1, m, q float64;

  fmt.Print("Inserisci la coordinata x del punto: ");
  fmt.Scan(&x1);
  
  fmt.Print("Inserisci la coordinata y del punto: ");
  fmt.Scan(&y1);

  fmt.Print("Inserisci il coefficiente angolare m: ");
  fmt.Scan(&m);
  
  fmt.Print("Inserisci il termine noto q: ");
  fmt.Scan(&q);

  retta := m * x1 + q;

  if math.Abs(y1 - retta) <= EPSILON {
    fmt.Println("sulla retta");
  } else if y1 > retta {
    fmt.Println("sopra");
  } else {
    fmt.Println("sotto");
  }
}
