// A condition can be either true or false.
//
// Go supports the usual comparison operators from mathematics:
//
// Less than <
// Less than or equal <=
// Greater than >
// Greater than or equal >=
// Equal to ==
// Not equal to !=
// Additionally, Go supports the usual logical operators:
//
// Logical AND &&
// Logical OR ||
// Logical NOT !

// Example
// x > y
// x != y
// (x > y) && (y > z)
// (x == y) || z

// Go has the following conditional statements:
//
// Use if to specify a block of code to be executed, if a specified condition is true
// Use else to specify a block of code to be executed, if the same condition is false
// Use else if to specify a new condition to test, if the first condition is false
// Use switch to specify many alternative blocks of code to be executed

package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}

func exampleOne() {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
}

func exampleTwo() {
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}

func exampleThree() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}

// The brackets in the else statement should be like } else {
// Having the else brackets in a different line will raise an error

func exampleFour() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}

func examleFive() {
  a := 14
  b := 14
  if a < b {
    fmt.Println("a is less than b.")
  } else if a > b {
    fmt.Println("a is more than b.")
  } else {
    fmt.Println("a and b are equal.")
  }
}

func exampleSix() {
  x := 30
  if x >= 10 {
    fmt.Println("x is larger than or equal to 10.")
  else if x > 20
    fmt.Println("x is larger than 20.")
  } else {
    fmt.Println("x is less than 10.")
  }
}

// nested if statement

func exampleSeven() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
