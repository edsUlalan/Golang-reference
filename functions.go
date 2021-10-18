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
