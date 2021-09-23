package main
import ("fmt")

// print the values of i and j
func main() {
  var i, j string = "Hello", "World"

  fmt.Print(i)
  fmt.Print(j)
}

// print the arguments in new lines
func myFunction() {
  var i, j string = "World", "Hello"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}

// it is also possible to only use one Print() for printing multiple variables
func myFunctionTwo() {
  var i, j string = "First", "Name"

  fmt.Print(i, "\n", j)
}

// if we want to add a space  between string arguments, we need to use " "
func myFunctionThree() {
  var i, j string = "Second", "String"

  fmt.Print(i, " ", j)
}

//  print() inserts a space between the arguments if neither are strings

func myFunctionFour() {
  var i, j = 10, 20

  fmt.Print(i, j)
}

// The Println() function is similar to Print() with the difference
// that a whitespace is added between the arguments, and a newline is added at the end

func myFunctionFive() {
  var i, j string = "Third", "String"

  fmt.Println(i,j)
}

// The Printf() function first formats its argument based on the given formatting verb and then prints them
// %v is used to print the value of the arguments
// %T is used to print the type of the arguments

func myFunctionFourth() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
