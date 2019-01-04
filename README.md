# Golang Programming Language
## Hello Word  
Output:   
```go
Hello World
```

---

## Variables  
- string  
- bool  
- int, int8, int16, int32, int64  
- uint, uint8, uint16, uint32, uint64, uintptr  
- byte - alias for uint8  
- rune -alias for int32  
- float32, float64  
- complex64, complex128  

Output:  
```go
Henry duck 37 false 1.3 Zion Zion@gmail.com
float64
```

---  

## Packages  
Import:
```go
"fmt"
"math"

"github.com/henry226/go_crash_course/03_packages/strutil" (self made package)
``` 
Output:
```go
math.Floor(2.7) =  2
math.Ceil(2.7) =  3
math.Sqrt(16) =  4
Reverse 'hello' =  olleh
```

---

## Function
Functions:
```go
func greeting(name string) string {
	return "Hello " + "name"
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
```
Output:
```go
greeting('Mary') = Hello name
getSum(10,11) =  21 
```

---

## Arrays_Slices
Array and Slice:
```go
// Arrays
var fruitArr [2]string

// Assign values
fruitArr[0] = "Banana"
fruitArr[1] = "Mango"  

// Declare and assign
carArr := [2]string{"Honda", "Toyota"}  

// slice (array with dynamic size)
foodSlice := []string{"Sushi", "Noodle", "Steak", "Burger"}
```
Output:
```go
fruitArr:  [Banana Mango]
fruitArr[0]:  Banana
carArr:  [Honda Toyota]
carArr[0]:  Honda
foodSlice:  [Sushi Noodle Steak Burger]
len(foodSlice):  4
foodSlice[1:3] [Noodle Steak]
```

---

## Conditionals
If-else condition:
```go
if x < y {
	fmt.Printf("%d is less than %d\n", x, y)
} else if x == y {
	fmt.Printf("%d is equal to %d\n", x, y)
} else {
	fmt.Printf("%d is greater than %d\n", x, y)
}
```
Switch Statement:
```go
switch color {
case "red":
	fmt.Println("color is red")
case "blue":
	fmt.Println("color is blue")
default:
	fmt.Println("color is not blue nor red")
}
```
Output:
```go
25 is equal to 25
color is not blue nor red
```

--- 

## Loops
Long method:
```go
i := 1
for i <= 10 {
	fmt.Print(i, " ")
	i++
}
```
Short method:
```go
for i := 1; i <= 10; i++ {
	fmt.Printf("%d ", i)
}
```
FizzBuzz Game: 
```go
counter := 1
for i := 1; i <= 100; i++ {
	if i%15 == 0 {
		fmt.Print("FizzBuzz ")
	} else if i%3 == 0 {
		fmt.Print("Fizz ")
	} else if i%5 == 0 {
		fmt.Print("Buzz ")
	} else {
		fmt.Print(i, " ")
	}

	if counter%10 == 0 {
		fmt.Println()
	}
	counter++
}
```
Output:
```go
1 2 3 4 5 6 7 8 9 10
1 2 3 4 5 6 7 8 9 10
FizzBuzz Game:
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz
11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz
Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz
41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz
Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz
61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz
71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz
Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz
91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```

---

## Maps  
Declare Map:
```go
// Define map
emails := make(map[string]string)

// Declare map and add key value
students := map[int]string{1: "Mary", 2: "Anita", 3: "Andy"}
```
Assign key value:
```go
// Assign Key Value
emails["Peter"] = "peter@email.com"
emails["Dick"] = "dick@email.com"
emails["Emma"] = "emma@email.com"

// Declare map and add key value
students := map[int]string{1: "Mary", 2: "Anita", 3: "Andy"}
```
Output key's value:
```go
fmt.Println("Peter's email:", emails["Peter"])
```
Delete key:
```go
delete(emails, "Dick")
```
Map length: 
```go 
len(emails)
```
Output: 
```go
All emails: map[Peter:peter@email.com Dick:dick@email.com Emma:emma@email.com]
Total 3 email(s)
Peter's email: peter@email.com
Delete Dick's email
All emails: map[Peter:peter@email.com Emma:emma@email.com]
Student's map: map[2:Anita 3:Andy 1:Mary]
```

---

## Range
Declare slice and map:
```go
// Slice
ids := []int{33, 76, 54, 23, 11, 22}
// Map
students := map[int]string{1: "Mary", 2: "Anita", 3: "Andy"}
```
Loop through slice:
```go
for i, id := range ids {
	fmt.Printf("%d - ID: %d\n", i, id)
}
```
Not using index:
```go
for _, id := range ids {
	fmt.Printf("IDs: %d\n", id)
}
```
Add all elements in slice:
```go
sum := 0
for _, id := range ids {
	sum += id
}
fmt.Println("Sum of all ids:", sum)
```
Range with map:
```go
for key, value := range students {
	fmt.Printf("%d : %s\n", key, value)
}
```
Range with map without value:
```go
for key := range students {
	fmt.Printf("Keys: %d\n", key)
}
```
Output:
```go
0 - ID: 33
1 - ID: 76
2 - ID: 54
3 - ID: 23
4 - ID: 11
5 - ID: 22
IDs: 33
IDs: 76
IDs: 54
IDs: 23
IDs: 11
IDs: 22
Sum of all ids: 219
1 : Mary
2 : Anita
3 : Andy
Keys: 1
Keys: 2
Keys: 3
```

---

## Pointers
Declare variables:
```go
a := 5
b := &a
```
Print the types of a and b:
```go
fmt.Println(a, b)
fmt.Printf("a = %T, b = %T\n", a, b)
```
Use * to read val from address:
```go
fmt.Println("*&a = ", *&a, "*b = ", *b)
```
Change val with pointer: 
```go
*b = 10
fmt.Println("a = ", a)
```
Output: 
```go
5 0xc00005a058
a = int, b = *int
*&a = 5 *b = 5
a =  10
```

---

## Closures
Function with anonymous functions:
```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```
Main:
```go
sum := adder()
for i := 0; i < 10; i++ {
	fmt.Println(sum(i))
}
```
Output:
```go
0
1
36
10
15
21
28
36
45
```

---

## Structs
Define Person struct: 
```go
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// Short cut
	// firstName, lastName, city, gender string
	// age int
}
```
Value reciever method
```go
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + " years old."
}
```
Pointer reciever method
```go
func (person *Person) hasBirthday() {
	fmt.Println("+1 age (called person1.hasBirthday())")
	person.age++
}

func (person *Person) getMarried(spouseLastName string) {
	fmt.Println(person.firstName, person.lastName, "is married to", spouseLastName)

	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}
```
Init person using struct: 
```go 
person1 := Person{firstName: "Mary", lastName: "Wong", city: "Washington", gender: "f", age: 30}
person2 := Person{"Peter", "Chang", "Franklin", "m", 44}
```
Print and modify person info:
```go
// Output Person
fmt.Println("person1:", person1, "\nperson2:", person2)
fmt.Println("person1 firstname:", person1.firstName)

// Change person info
person1.age = person1.age + 10
fmt.Println("person1:", person1, "(age modified)")

// Output greeting method (value reciever)
fmt.Println(person1.greet())

// Output hasBirthday method (pointer reciever)
person1.hasBirthday()
person1.hasBirthday()
fmt.Println(person1.greet())

// Output getMarried method (pointer reciever)
person1.getMarried("John Li")
fmt.Println(person1.greet())
person2.getMarried("Yal Chen")
fmt.Println(person2.greet(), "(last name no change because Peter is male.)")
```
Output: 
```go
person1: {Mary Wong Washington f 30}
person2: {Peter Chang Franklin m 44}
person1 firstname: Mary
person1: {Mary Wong Washington f 40} (age modified)
Hello, my name is Mary Wong and I am 40 years old.
+1 age (called person1.hasBirthday())
+1 age (called person1.hasBirthday())
Hello, my name is Mary Wong and I am 42 years old.
Mary Wong is married to John Li
Hello, my name is Mary John Li and I am 42 years old.
Peter Chang is married to Yal Chen
Hello, my name is Peter Chang and I am 44 years old. (last name no change because Peter is male.)
```

---

## Interfaces
Define interface:
```go
type Shape interface {
	area() float64
}
```
Define structs:
```go
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}
```
Define methods:
```go
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}
```
Init circle and rectangle structs:
```go 
circle := Circle{x: 0, y: 0, radius: 5}
rect := Rectangle{width: 12, height: 5}
```
Print areas:
```go
fmt.Printf("Circle Area: %.2f\n", getArea(circle))
fmt.Printf("Rectangle Area: %.2f\n", getArea(rect))
```
Output:
```go
Circle Area: 78.54
Rectangle Area: 60.00
```

---

## Web
Functions for website content:
```go
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index Page</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
}
```
Connect to localhost:
```go
http.HandleFunc("/", index)
http.HandleFunc("/about", about)
fmt.Println("Server Starting... localhost:3000")
http.ListenAndServe(":3000", nil)
```
Output:
```go
Server Starting... localhost:3000
```
