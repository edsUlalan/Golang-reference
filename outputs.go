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

// go formatting verbs

// general
// Verb	Description
// %v	  Prints the value in the default format
// %#v	  Prints the value in Go-syntax format
// %T	  Prints the type of the value
// %%	  Prints the % sign

func verbs() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i) // 15.5
  fmt.Printf("%#v\n", i) // 15.5
  fmt.Printf("%v%%\n", i) // 15.5%
  fmt.Printf("%T\n", i) // float64

  fmt.Printf("%v\n", txt) // Hello World!
  fmt.Printf("%#v\n", txt) // "Hello World!"
  fmt.Printf("%T\n", txt) // string
}

// integer formatting verbs
// Verb	Description
// %b	  Base 2
// %d	  Base 10
// %+d	  Base 10 and always show sign
// %o	  Base 8
// %O	  Base 8, with leading 0o
// %x	  Base 16, lowercase
// %X	  Base 16, uppercase
// %#x	  Base 16, with leading 0x
// %4d	  Pad with spaces (width 4, right justified)
// %-4d	Pad with spaces (width 4, left justified)
// %04d	Pad with zeroes (width 4)

func integers() {
  var i = 15

  fmt.Printf("%b\n", i) // 1111
  fmt.Printf("%d\n", i) // 15
  fmt.Printf("%+d\n", i) // +15
  fmt.Printf("%o\n", i) // 17
  fmt.Printf("%O\n", i) // 0o17
  fmt.Printf("%x\n", i) // f
  fmt.Printf("%X\n", i) // F
  fmt.Printf("%#x\n", i) // 0xf
  fmt.Printf("%4d\n", i) // 15
  fmt.Printf("%-4d\n", i) // 15
  fmt.Printf("%04d\n", i) // 0015
}

// String Formatting Verbs
// Verb	Description
// %s	  Prints the value as plain string
// %q	  Prints the value as a double-quoted string
// %8s	  Prints the value as plain string (width 8, right justified)
// %-8s	Prints the value as plain string (width 8, left justified)
// %x	  Prints the value as hex dump of byte values
// % x	  Prints the value as hex dump with spaces

func stringVerb() {
  var txt = "Hello"

  fmt.Printf("%s\n", txt) // Hello
  fmt.Printf("%q\n", txt) // "Hello"
  fmt.Printf("%8s\n", txt) //    Hello
  fmt.Printf("%-8s\n", txt) // Hello
  fmt.Printf("%x\n", txt) // 48656c6c6f
  fmt.Printf("% x\n", txt) // 48 65 6c 6c 6f
}

// boolean
// Verb	Description
// %t	  Value of the boolean operator in true or false format (same as using %v)

func booleanverb() {
  var i = true
  var j = false

  fmt.Printf("%t\n", i) // true
  fmt.Printf("%t\n", j) // false
}

// float Formatting
// Verb	Description
// %e	  Scientific notation with 'e' as exponent
// %f	  Decimal point, no exponent
// %.2f	Default width, precision 2
// %6.2f	Width 6, precision 2
// %g	  Exponent as needed, only necessary digits

func floatVerb() {
  var i = 3.141

  fmt.Printf("%e\n", i) // 3.141000e+00
  fmt.Printf("%f\n", i) // 3.141000
  fmt.Printf("%.2f\n", i) // 3.14
  fmt.Printf("%6.2f\n", i) //   3.14
  fmt.Printf("%g\n", i) // 3.141
}
