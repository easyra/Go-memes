package main

import "fmt"

const (
  Pi = 3.14
  Truth= false
  Big = 1<< 62
  small = Big >> 61
)

func main(){
  const Greeting = "Hi"
  fmt.Println(Greeting)
  fmt.Println(Pi)
  fmt.Println(Truth)
	fmt.Println(Big)
	name := "Caperica-Six"
  aka := fmt.Sprintf("Number %d", 6)
  fmt.Printf("%s is also known as %s",
    name,aka)
}