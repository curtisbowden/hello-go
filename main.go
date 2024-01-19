package main

import (
  "fmt"
  "math/rand"
  "time"
  "github.com/ctomkow/go-hello"
)
// go get github.com/ctomkow/go-hello

func main() {
  rand.Seed(time.Now().UnixNano())
  gohelo.GoHello()
  fmt.Println("Hello Go!")
  fmt.Println("Go can print random numbers:", rand.Intn(100))
  fmt.Printf("")
}
