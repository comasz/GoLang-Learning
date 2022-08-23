# What is Golang?
Go borrows static typing and run-time efficiency from **C** and readability and usability from **Python** and **JavaScript** with its main goal of stability at its core. Some key features in go: 
- Structurally typed: user-defined data type that allow you to combine data items of different kinds. implies a given value must be compatible to a defined structure.
- Garbage collection: reclaims memory that was allocated by program but no longer referenced.
- Compiled: translates source code to machine code binary.
- Concurrency: when a program has two or more tasks that run individually of each other but at the same time. (example: [Goroutines]())
- Recursion: function that calls itself. Runs the risk of infinate loop.


## Common Terminology
- Abstraction: hides complexity from the user. 
- Structure: collection of data fields with declared data types.
- Token: is either a keyword, an identifier, a constant, a string literal, or a symbol.
- Decleration: when you assign type to variable name
- Initialization: when you assign value to variable
- Rune: string literal utf-8 sequence (example: `U+1F600` = 😀) 
- Call By Reference: when passing arguement into a function, copies the address of an argument into the parameter.
- Pointer: a variable whose value is a direct address in the memory. Used for call by reference. (operator = `*`)
- Ampersand: `&` operator which pulls an address from memory
- Hexadecimal: represents memory address


## Data Types
- Map: data type that maps unique keys to values
- Array: collection data of the same type, has a fixed size. 
- Slice: Dynamic sized array. you are not allowed to store different type of elements in the same slice.
- Interfaces: represents a set of method signatures.

### Go's Replacement of Class Inheritance
- embedding: automated form of composition
- interfaces: are named collections of method signatures.
- method signature:function name, input parameters and return type.

### Modeling Structured Data Such as Json & YAML
- The `interface{}` type can be used to model structured data by representing it as a `map[string]interface{}` (map of string to empty interface). This describes data in the form of a dictionary with string keys and values of any type.
- `interface{}` can refer to any value, it is a limited way to escape the restrictions of static typing.
	- convert it to a more useful type via a `type assertion` or `type switch`, or inspect it with Go's `reflect` package.
- Golang to Json structure -> [Link](https://mholt.github.io/json-to-go/)
### Array & Slice
- Array: data structure that holds a fixed size of **sequential** collection of elements of the same type.
- Slice: like an array except size is not fixed. 
```go
func main() {
	// Define number of elements
	var n = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// Omit size of array it will hold only the amount of values declared. known as slice
	var m = []float32{10500.0, 5.0, 7.4, 6.0}

	// output each array element's value
	for i := range n {
		fmt.Println(n[i])
	}
	for i := range m {
		fmt.Println(n[i])
	}
}
```
### Structures
A record of different objects with defined data types.
```go
package main
import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    // Declare Book1 of type Book
   var Book2 Books    // Declare Book2 of type Book
 
   // book 1 specification
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   // book 2 specification
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
   }
}
   ```
   
### Map 
Group of key-value pairs much like a dictionary.
```go
package main
import "fmt"
   
func main() {
	var countryCapitalMap map[string]string // declare map variable
	countryCapitalMap = make(map[string]string) //create a map

	// insert key-value pairs in the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// print map using keys
	for country := range countryCapitalMap {
	   fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["United States"] // test if entry is present in the map

	// if ok is true, entry is present otherwise entry is absent
	if(ok){
	   fmt.Println("Capital of United States is", capital)  
	} else {
	   fmt.Println("Capital of United States is not present") 
	}
	
	for i := range countryCapitalMap{
		fmt.Println(i)
	}
}
 ```
 note: Map keys are stored in random order. If you are looking to maintain order use `orderedmap` package instead.
 
### Pointers
Allows memory addresses to be referenced so referenced values arent copies to a new memory address. 
You must declare a pointer before you can use it to store any variable address
```go
package main
import "fmt"

func main() {
   var a int = 20   // actual variable declaration 
   var ip *int      // pointer variable declaration 

   ip = &a  // store address of a in pointer variable

   fmt.Printf("Address of a variable: %x\n", &a  )

   fmt.Printf("Address stored in ip variable: %x\n", ip ) // address stored in pointer variable 
   
   fmt.Printf("Value of *ip variable: %d\n", *ip ) // access the value using the pointer 
}
```

### Recursion
Below is an example of fibonacci sequence run as a recursive function.
```go
package main
import "fmt"

func fibonaci(i int) (ret int) {
   if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
   var i int
   for i = 0; i < 10; i++ {
      fmt.Printf("%d ", fibonaci(i))
   }
}
```

### Inheritance

``` go
package main
import "fmt"

// define an interface
type Car interface {
   engine() string
}

// define structures
type InternalCombustion struct {
   gas, horsepower, color string
}

type Electric struct {
   horsepower, color string
}

// implementation of c.engine()
func(ic InternalCombustion) engine() string {
   return ic.gas + " " + ic.horsepower + " " + ic.color
}

// implementation of c.engine()
func(e Electric) engine() string {
   return e.horsepower + " " + e.color
}

func getengine(c Car) string {
   return c.engine()
}

func main() {
   ic := InternalCombustion{gas:"deisel",horsepower:"250hp",color:"white"}
   e := Electric {horsepower:"120hp", color:"black"}
   
   fmt.Printf("Internal combustion engine: %v\n",getengine(ic))
   fmt.Printf("Electric engine: %v\n",getengine(e))
}
```

## Error Handling
Functions normally return error as last return value. Use errors.New to construct a basic error message

``` go
package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value), nil
}
func main() {
   result, err:= Sqrt(-1)
   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
   result, err = Sqrt(9)
   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
}
```

## Simple Syntax
```
Construct: m := map[key]value{}
Insert: m[k] = V
Lookup: V = m[k]
Delete: delete (m, k)
Iterate: for k, v := range m
Size: len(m)
```

