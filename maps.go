// Maps are used to store data values in key:value pairs.
//
// Each element in a map is a key:value pair.
//
// A map is an unordered and changeable collection that does not allow duplicates.
//
// The length of a map is the number of its elements. You can find it using the len() function.
//
// The default value of a map is nil.
//
// Maps hold references to an underlying hash table.
//
// Go has multiple ways for creating maps.

package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo":1, "Bergen":2, "Trondheim":3, "Starvanger":4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

// a   map[brand:Ford model:Mustang year:1964]
// b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]

// creating maps using make() function

func useMake() {
  var a = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"
                                 // a is no longer empty
  b := make(map[string]int)
  b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

// a   map[brand:Ford model:Mustang year:1964]
// b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]

// creating empty map

func emptyMap() {
  var a = make(map[string]string)
  var b map[string]string

  fmt.Println(a == nil)
  fmt.Println(b == nil)
}

// false
// true

// The map key can be of any data type for which the equality operator (==) is defined. These include:
//
// Booleans
// Numbers
// Strings
// Arrays
// Pointers
// Structs
// Interfaces (as long as the dynamic type supports equality)
// Invalid key types are:
//
// Slices
// Maps
// Functions
// These types are invalid because the equality operator (==) is not defined for them.

// The map values can be any type.

// accessing map elements

func accesMap() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}

// Ford

// updating and adding map elements

func updateMap() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  a["year"] = "1970" // Updating an element
  a["color"] = "red" // Adding an element

  fmt.Println(a)
}

// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford color:red model:Mustang year:1970]

// remove element from map

func deleteMap() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  delete(a,"year")

  fmt.Println(a)
}

// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford model:Mustang]

// check for specific elements in a map

func checkMap() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

  val1, ok1 := a["brand"] // Checking for existing key and its value
  val2, ok2 := a["color"] // Checking for non-existing key and its value
  val3, ok3 := a["day"]   // Checking for existing key and its value
  _, ok4 := a["model"]    // Only checking for existing key and not its value

  fmt.Println(val1, ok1)
  fmt.Println(val2, ok2)
  fmt.Println(val3, ok3)
  fmt.Println(ok4)
}

// Ford true
// false
// true
// true

// In this example, we checked for existence of different keys in the map.
//
// The key "color" does not exist in the map. So the value is an empty string ('').
//
// The ok2 variable is used to find out if the key exist or not.
// Because we would have got the same value if the value of the "color" key was empty.
// This is the case for val3.

// If two map variables refer to the same hash table,
// changing the content of one variable affect the content of the other.

func changeMap() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := a

  fmt.Println(a)
  fmt.Println(b)

  b["Year"] = "1970"
  fmt.Println("After change to b:")

  fmt.Println(a)
  fmt.Println(b)
}

// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford model:Mustang year:1964]
// After change to b:
// map[Year:1970 brand:Ford model:Mustang year:1964]
// map[Year:1970 brand:Ford model:Mustang year:1964]

// iterate over maps

func iterateMaps() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  for k, v := range a {
    fmt.Printf("%v : %v, ", k, v)
  }
}

// two : 2, three : 3, four : 4, one : 1,

// iterate maps over specific order

func specificMaps() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var b = []string             // defining the order
  b = append(b, "one", "two", "three", "four")

  for k, v := range a {        // loop with no order
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }
}

// two : 2, three : 3, four : 4, one : 1,
// one : 1, two : 2, three : 3, four : 4,
