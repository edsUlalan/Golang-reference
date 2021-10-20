// Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
//
// In the following example, testcount() is a function that calls itself.
// We use the x variable as the data, which increments with 1 (x + 1) every time we recurse.
// The recursion ends when the x variable equals to 11 (x == 11).

package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}

// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}

// 24
