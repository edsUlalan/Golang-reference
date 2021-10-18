// Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.
//
// In the example below, we create a function named "myMessage()".
// The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function.
// The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ()

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

// A function can be called multiple times.

func callTwice() {
  myMessage()
  myMessage()
}

// A function name must start with a letter
// A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
// Function names are case-sensitive
// A function name cannot contain spaces
// If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used
// Give the function a name that reflects what the function does!


// function with parameters

// The following example has a function with one parameter (fname) of type string.
// When the familyName() function is called, we also pass along a name (e.g. Liam),
// and the name is used inside the function, which outputs several different first names,
// but an equal last name

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func results() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}

// Hello Liam Refsnes
// Hello Jenny Refsnes
// Hello Anja Refsnes

// multiple parameters

func familyNameMultiple(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func resultsTwo() {
  familyNameMultiple("Liam", 3)
  familyNameMultiple("Jenny", 14)
  familyNameMultiple("Anja", 30)
}

// Hello 3 year old Liam Refsnes
// Hello 14 year old Jenny Refsnes
// Hello 30 year old Anja Refsnes

// function returns

func myFunction(x int, y int) int {
  return x + y
}

func resultsThree() {
  fmt.Println(myFunction(1, 2))
}

// 3

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func resultsFour() {
  fmt.Println(myFunction(1, 2))
}

// 3

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}

// store the return value in a parameter
