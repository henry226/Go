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