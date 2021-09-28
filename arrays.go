package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1) // [1 2 3]
  fmt.Println(arr2) // [4 5 6 7 8]
}

// syntax

var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred

array_name := [length]datatype{values} // here length is defined
or
array_name := [...]datatype{values} // here length is inferred

// declaring two array with inferred lengths

func arraySample() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1) // [1 2 3]
  fmt.Println(arr2) // [4 5 6 7 8]
}

// declare array of string

func arrayString() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars) // [Volvo BMW Ford Mazda]
}

// access elements of of an array

func element() {
  prices := [3]int{10, 20, 30}

  fmt.Println(prices[0]) // 10
  fmt.Println(prices[2]) // 30
}

// change value of a specific array element

func change() {
  prices := [3]int{10, 20, 30}

  prices[2] = 50
  fmt.Println(prices) // [10 20 50]
}

// array initialization

 func initialize() {
   arr1 := [5]int{} //not initialized
   arr2 := [5]int{1,2} //partially initialized
   arr3 := [5]int{1,2,3,4,5} //fully initialized

   fmt.Println(arr1) // [0 0 0 0 0]
   fmt.Println(arr2) // [1 2 0 0 0]
   fmt.Println(arr3) // [1 2 3 4 5]
 }

 // initialize specific elements

 func specific() {
   arr1 := [5]int{1:10,2:40}

   fmt.Println(arr1) // [0 10 40 0 0]
 }

// 1:10 means: assign 10 to array index 1 (second element).
// 2:40 means: assign 40 to array index 2 (third element).

// find length of an array

func arrayLength() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1)) // 4
  fmt.Println(len(arr2)) // 6
}
