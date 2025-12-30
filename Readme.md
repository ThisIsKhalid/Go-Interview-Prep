Below is a **complete, interview-ready, project-oriented breakdown** of **each topic**.

You can:

* use it as **revision notes**
* prepare for **Go interviews**
* directly **apply concepts in real projects**

---

````md
# Go Language – Complete Breakdown (Senior-Level Notes)

> These notes are written for:
> - Learning Go deeply
> - Revising core concepts
> - Interview preparation
> - Building real-world Go projects

---

## 1. Variables and Data Types

### Variables
Go is a **statically typed** language.

#### Declaration styles
```go
var a int = 10
var b = 20        // type inferred
c := 30           // short declaration (inside functions only)
````

#### Multiple variables

```go
var x, y int = 1, 2

var (
	name string = "Go"
	age  int    = 10
)
```

---

### Data Types

#### Basic Types

```go
int, int8, int16, int32, int64
uint, uint8 (byte), uint16, uint32, uint64
float32, float64
bool
string
```

#### Composite Types

```go
array
slice
map
struct
```

#### Reference Types

```go
slice
map
channel
pointer
function
```

---

## 2. If-Else and Switch Case

### if-else

```go
age := 18

if age >= 18 {
	fmt.Println("Adult")
} else {
	fmt.Println("Minor")
}
```

#### if with short statement

```go
if score := 85; score >= 80 {
	fmt.Println("A Grade")
}
```

---

### switch

```go
day := "Monday"

switch day {
case "Saturday", "Sunday":
	fmt.Println("Weekend")
default:
	fmt.Println("Working day")
}
```

> Go switch **does not fall through by default**

---

## 3. Functions with Parameters and Return Types

### Basic function

```go
func add(a int, b int) int {
	return a + b
}
```

### Multiple return values

```go
func divide(a, b int) (int, int) {
	return a / b, a % b
}
```

### Named return

```go
func sum(a, b int) (result int) {
	result = a + b
	return
}
```

---

## 4. Scope (Very Important)

Scope defines **where a variable is accessible**.

---

## 5. Local Scope and Block Scope

### Local Scope

```go
func test() {
	x := 10
	fmt.Println(x)
}
```

### Block Scope

```go
if true {
	y := 20
	fmt.Println(y)
}
```

> `y` is NOT accessible outside the block

---

## 6. Package Scope

```go
package main

var appName = "MyApp" // package-level variable

func main() {
	fmt.Println(appName)
}
```

> Accessible inside the same package
> Exported if name starts with **capital letter**

---

## 7. Variable Shadowing

When a variable in inner scope **overrides** outer variable.

```go
x := 10

if true {
	x := 20 // shadows outer x
	fmt.Println(x) // 20
}

fmt.Println(x) // 10
```

⚠️ Common interview trap

---

## 8. Function Types & Named Functions

### Named Function

```go
func greet() {
	fmt.Println("Hello")
}
```

### Function Type

```go
type MathFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

var f MathFunc = add
```

---

## 9. init Function & Function Expression

### init()

* Runs **before main**
* Used for setup

```go
func init() {
	fmt.Println("Init executed")
}
```

### Function Expression

```go
add := func(a, b int) int {
	return a + b
}

fmt.Println(add(2, 3))
```

---

## 10. Anonymous Function and IIFE

### Anonymous Function

```go
func() {
	fmt.Println("Anonymous")
}()
```

### IIFE (Immediately Invoked Function Expression)

```go
func(x int) {
	fmt.Println(x)
}(10)
```

Used in:

* goroutines
* closures
* quick logic

---

## 11. A “Noob” Function (Simple Example)

```go
func sayHello() {
	fmt.Println("Hello Go")
}
```

> Every Go developer starts here 🙂

---

## 12. Parameters vs Arguments

```go
func add(a int, b int) int { // parameters
	return a + b
}

add(2, 3) // arguments
```

---

## 13. First Order Function

A function that:

* does NOT accept another function
* does NOT return a function

```go
func square(x int) int {
	return x * x
}
```

---

## 14. Higher Order Function

A function that:

* accepts a function OR
* returns a function

```go
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
```

---

## 15. First Class Function

Functions in Go are **first-class citizens**:

* assign to variable
* pass as argument
* return from function

```go
f := func() {
	fmt.Println("First class function")
}

f()
```

---

## 16. Callback Function

A function passed into another function and executed later.

```go
func process(callback func()) {
	callback()
}

process(func() {
	fmt.Println("Callback executed")
})
```

Used heavily in:

* concurrency
* async workflows
* libraries

---

## 17. Go Internal Memory Model

### Memory Segments

1. **Code Segment**

   * Compiled instructions

2. **Data Segment (Global Memory)**

   * Global & package variables

3. **Stack**

   * Function calls
   * Local variables
   * Fast allocation

4. **Heap**

   * Dynamically allocated memory
   * Managed by GC

5. **Garbage Collector (GC)**

   * Automatically frees unused heap memory

---

## 18. Stack Frame

Each function call creates a **stack frame** containing:

* parameters
* local variables
* return address

```go
func A() {
	B()
}

func B() {
	fmt.Println("Stack Frame Example")
}
```

Stack grows & shrinks automatically.

---

## 19. Two Phases in Go Program

### 1. Compilation Phase

* Syntax check
* Type check
* Build binary

### 2. Execution Phase

* Program runs
* Memory allocated
* Functions executed

---

## 20. go run vs go build

### go run

```bash
go run main.go
```

Steps:

```
compile → temporary binary → execute → delete
```

Used for:

* development
* quick testing

---

### go build

```bash
go build main.go
```

Steps:

```
compile → create ./main binary
```

Used for:

* production
* deployment