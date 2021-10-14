// statement1 Initializes the loop counter value.
// statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
// statement3 Increases the loop counter value.

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

// 0
// 1
// 2
// 3
// 4

// i:=0; - Initialize the loop counter (i), and set the start value to 0
// i < 5; - Continue the loop as long as i is less than 5
// i++ - Increase the loop counter value by 1 for each iteration
