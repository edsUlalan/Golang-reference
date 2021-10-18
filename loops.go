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

func countByTen() {
  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }
}

// 0
// 10
// 20
// 30
// 40
// 50
// 60
// 70
// 80
// 90
// 100

// i:=0; - Initialize the loop counter (i), and set the start value to 0
// i <= 100; - Continue the loop as long as i is less than or equal to 100
// i+=10 - Increase the loop counter value by 10 for each iteration

// continue statement

func skipValue() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}

// 0
// 1
// 2
// 4

// break statement

func breakValue() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}

// 0
// 1
// 2

// nested loop

func nestedLoop() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}

// big apple
// big orange
// big banana
// tasty apple
// tasty orange
// tasty banana

// range keyword
// The range keyword is used to more easily iterate over an array, slice or map.
// It returns both the index and the value

func rangeExample() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}

// 0      apple
// 1      orange
// 2      banana

// This example uses range to iterate over an array and print both the indexes and
// the values at each (idx stores the index, val stores the value)

func rangeExampleTwo() {
  fruits := [3]string{"apple", "orange", "banana"}
  for _, val := range fruits {
     fmt.Printf("%v\n", val)
  }
}

// To only show the value or the index, you can omit the other output using an underscore (_)

// apple
// orange
// banana

func rangeExampleThree() {
  fruits := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }
}

// Here, we want to omit the values (idx stores the index, val stores the value)

// 0
// 1
// 2
