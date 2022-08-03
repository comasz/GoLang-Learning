# What is Golang?
Go borrows static typing and run-time efficiency from **C** and readability and usability from **Python** and **JavaScript** with its main goal of stability at its core. Some key features in go: 
- Structurally typed: user-defined data type that allow you to combine data items of different kinds. implies a given value must be compatible to a defined structure.
- garbage collection: reclaims memory that was allocated by program but no longer referenced.
- compiled: translates source code to machine code binary.
- concurrency: 




## Common Terminology
- Abstraction: hides complexity from the user. 
- Structure: 
- Token: is either a keyword, an identifier, a constant, a string literal, or a symbol.

### Go's Replacement of Class Inheritance
- embedding: automated form of composition
- interfaces: 

### Modeling Structured Data Such as Json & YAML
- The `interface{}` type can be used to model structured data by representing it as a `map[string]interface{}` (map of string to empty interface). This describes data in the form of a dictionary with string keys and values of any type.
- `interface{}` can refer to any value, it is a limited way to escape the restrictions of static typing.
    - convert it to a more useful type via a `type assertion` or `type switch`, or inspect it with Go's `reflect` package.

### Array
Data structure that holds a fixed size of **sequential** collection of elements of the same type.
```go
func main() {
	// Define number of elements
	var n = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// Omit size of array it will hold only the amount of values declared
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
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    /* Declare Book1 of type Book */
   var Book2 Books    /* Declare Book2 of type Book */
 
   /* book 1 specification */
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   /* book 2 specification */
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
   ```
