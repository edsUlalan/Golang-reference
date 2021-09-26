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

// int	Depends on platform:
//     32 bits in 32 bit
//     systems and 64 bit
//     in 64 bit systems	- 2147483648 to 2147483647 in 32 bit systems and
//                        - 9223372036854775808 to 9223372036854775807 in 64 bit systems
// int8	  8 bits/1 byte	  - 128 to 127
// int16	16 bits/2 byte	- 32768 to 32767
// int32	32 bits/4 byte	- 2147483648 to 2147483647
// int64	64 bits/8 byte	- 9223372036854775808 to 9223372036854775807

// unsigned

func unsignedInteger() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

// uint	Depends on platform:
//       32 bits in 32 bit
//       systems and 64 bit
//       in 64 bit systems	  0 to 4294967295 in 32 bit systems and
//                           0 to 18446744073709551615 in 64 bit systems
// uint8	  8 bits/1 byte	    0 to 255
// uint16	16 bits/2 byte	  0 to 65535
// uint32	32 bits/4 byte	  0 to 4294967295
// uint64	64 bits/8 byte	  0 to 18446744073709551615

// float data type

func floatExample() {
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

// float64 -  can store a larger set of numbers than float32

func floatExampleTwo() {
  var x float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", x, x)
}

// float32	32 bits	-3.4e+38 to 3.4e+38.
// float64	64 bits	-1.7e+308 to +1.7e+308.

// string data type

func stringExample() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %\n", txt1, txt1)  // Type: string, value: Hello!
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2) // Type: string, value:
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3) // Type: string, value: World 1
}
