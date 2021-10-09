// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.
// In Go, there are several ways to create a slice:

// Using the []datatype{values} format
// Create a slice from an array
// Using the make() function

// syntax
slice_name := []datatype{values}
myslice := []int // empty slice of 0 length and 0 capacity
myslice := []int{1,2,3} // a slice of integers of length 3 and also the capacity of 3

// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1)) // 0
  fmt.Println(cap(myslice1)) // 0
  fmt.Println(myslice1) // []

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2)) // 4
  fmt.Println(cap(myslice2)) // 4
  fmt.Println(myslice2) // [Go SLices Are Powerful]
}

// create a slice from an array
// syntax
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array

func exampleSlice() {
  arr1 := [6]int{10, 11, 12, 13, 14, 15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice) // myslice = [12 13]
  fmt.Printf("length = %d\n", len(myslice)) // length = 2
  fmt.Printf("capacity = %d\n", cap(myslice)) // capacity = 4
}

// create a slice with the make() function
// syntax
slice_name := make([]type, length, capacity)
// Note: If the capacity parameter is not defined, it will be equal to length.

func exampleSliceTwo() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [0 0 0 0 0]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 5
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 10

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2) // myslice2 = [0 0 0 0 0]
  fmt.Printf("length = %d\n", len(myslice2)) // length = 5
  fmt.Printf("capacity = %d\n", cap(myslice2)) // capacity = 5
}

// access elements of a slice

func exampleSliceThree() {
  prices := []int{10,20,30}

  fmt.Println(prices[0]) // 10
  fmt.Println(prices[2]) // 20
}

// change elements of a slice

func exampleSliceFour() {
  prices := []int{10,20,30}
  prices[2] = 50
  fmt.Println(prices[0]) // 10
  fmt.Println(prices[2]) // 50
}

// append elements to a slice
// syntax
slice_name = append(slice_name, element1, element2, ...)

func exampleSliceFive() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [1 2 3 4 5 6]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 6
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 6

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [1 2 3 4 5 6 20 21]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 8
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 12
}

// append one slice to another slice
// syntax
slice3 = append(slice1, slice2...)

func exampleSliceSix() {
  myslice1 := []int{1,2,3}
  myslice2 := []int{4,5,6}
  myslice3 := append(myslice1, myslice2...)

  fmt.Printf("myslice3=%v\n", myslice3) // myslice3=[1 2 3 4 5 6]
  fmt.Printf("length=%d\n", len(myslice3)) // length=6
  fmt.Printf("capacity=%d\n", cap(myslice3)) // capacity=6
}

// change the length of a slice

func exampleSliceSeven() {
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice1 := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [10 11 12 13]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 4
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 5

  myslice1 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [10 11]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 2
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 5

  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [10 11 20 21 22 23]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 6
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 10
}


// memory efficiency
// The copy() function creates a new underlying array with only the required elements for the slice.
// This will reduce the memory used for the program.
// syntax
copy(dest, src)

func exampleSliceEight() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers) // numbers = [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
  fmt.Printf("length = %d\n", len(numbers)) // length = 15
  fmt.Printf("capacity = %d\n", cap(numbers)) // capacity = 15

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy) // numbersCopy = [1 2 3 4 5]
  fmt.Printf("length = %d\n", len(numbersCopy)) // length = 5
  fmt.Printf("capacity = %d\n", cap(numbersCopy)) // capacity = 5
}
