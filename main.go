package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.seed(time.Now())
  fmt.Println("Hello Go!")
  fmt.Println("Go can print random numbers:", rand.Intn(10))
}
