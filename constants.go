// declaring constant using const keyword

package main
import ("fmt")

const PI = 3.14 // written in uppercase only

func main() {
  fmt.Println(PI)
}

// typed constants

const A int = 1

func mainTwo() {
  fmt.Println(A)
}

// untyped constants

const B = 2

func main() {
  fmt.Println(A)
}

// When a constant is declared, it is not possible to change the value later

// multiple constants declaration

const (
  A int = 1
  B = 3.14
  c = "Hi!"
)

func printMe() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
