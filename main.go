package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Hello Go!")
  fmt.Println("Go can print random numbers:", rand.Intn(100))
  fmt.Printf("")
}
