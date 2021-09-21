package main
import("fmt")

// var keyword

var variablename type = value // You always have to specify either type or value (or both)

// := sign

variablename := value

// In this case, the type of the variable is inferred from the value
// (means that the compiler decides the type of the variable, based on the value)
// It is not possible to declare a variable using :=, without assigning a value to it

// variable declaration with initial value

func main() {
  var student1 string = "John" // type is string
  var student2 = "Jane" // type is inferred
  x := 2 // type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}

// variable declaration without initial value

func mainTwo() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

// These variables are declared but they have not been assigned initial values.
// By running the code, we can see that they already have the default values of their respective types:
// a is ""
// b is 0
// c is false

// value assignment after declaration

func mainThree() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}

//  Difference Between var and :=

// var
// Can be used inside and outside of functions
// Variable declaration and value assignment can be done separately

// :=
// Can only be used inside functions
// Variable declaration and value assignment cannot be done separately (must be done in the same line)

var a int
var b int = 2
var c = 3

func mainFour() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
