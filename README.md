# Golang Programming Language
### Hello Word  
Output:   
```go
Hello World
```

---

### Variables  
string  
bool  
int, int8, int16, int32, int64  
uint, uint8, uint16, uint32, uint64, uintptr  
byte - alias for uint8  
rune -alias for int32  
float32, float64  
complex64, complex128  
<br>
Output:  
```go
Henry duck 37 false 1.3 Zion Zion@gmail.com
float64
```

---  

### Packages  
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

### Function
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

### Arrays_Slices
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

### Conditionals
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

### Loops
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

### Maps  
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