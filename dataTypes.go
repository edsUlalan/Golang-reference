package main
import ("fmt")

func main() {
  var a bool = true // Boolean
  var b int = 5 // Integer
  var c float32 = 3.14 // Floatiing point number
  var d string = "Hi!" // string

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float: ", c)
  fmt.Println("String: ", d)
}

// boolean data type

func myFunctionBoolean() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(bo4) // Returns true
}

// integer data type
// Signed integers - can store both positive and negative values
// Unsigned integers - can only store non-negative values

// signed

func signedInteger() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

// unsigned

func unsignedInteger() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
