# 📘 Go Language – Complete Fundamentals Textbook

---

## 1️⃣ Variables and Data Types

### 🔹 Variables

A **variable** stores a value in memory.

### Syntax

```go
var name type = value
```

### Examples

```go
var age int = 25
var name string = "Khalid"
```

### Short Declaration (Most Used)

```go
age := 25
name := "Khalid"
```

> `:=` **can only be used inside functions**

---

### 🔹 Data Types

#### Basic Types

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
bool
string
```

#### Example

```go
var price float64 = 99.99
var isActive bool = true
```

---

### Zero Values

If not initialized:

```go
int     -> 0
float   -> 0.0
bool    -> false
string  -> ""
```

---

## 2️⃣ If-Else and Switch Case

### 🔹 If-Else

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### If with Initialization

```go
if score := 80; score > 50 {
    fmt.Println("Pass")
}
```

---

### 🔹 Switch Case

```go
switch day {
case "Mon":
    fmt.Println("Monday")
case "Tue":
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

### Switch without Expression

```go
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18:
    fmt.Println("Adult")
}
```

---

## 3️⃣ Functions (Parameters & Return Types)

### 🔹 Function Definition

```go
func add(a int, b int) int {
    return a + b
}
```

### Multiple Return Values

```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}
```

### Named Return

```go
func sum(a, b int) (result int) {
    result = a + b
    return
}
```

---

## 4️⃣ Scope (Very Important)

### 🔹 Scope

Scope defines **where a variable is accessible**.

---

## 5️⃣ Local Scope & Block Scope

### Local Scope

```go
func test() {
    x := 10 // accessible only inside test()
}
```

### Block Scope

```go
if true {
    y := 20 // accessible only inside this block
}
```

---

## 6️⃣ Package Scope

Variables declared **outside functions** belong to package scope.

```go
package main

var appName = "Go App"

func main() {
    fmt.Println(appName)
}
```

Accessible **anywhere in the same package**.

---

## 7️⃣ Variable Shadowing

When an inner scope variable has **same name** as outer scope.

```go
x := 10

if true {
    x := 20 // shadows outer x
    fmt.Println(x) // 20
}

fmt.Println(x) // 10
```

⚠️ Shadowing can cause bugs if misunderstood.

---

## 8️⃣ Function Types

Functions have **types**.

```go
type MathFunc func(int, int) int
```

### Usage

```go
func add(a, b int) int {
    return a + b
}

var f MathFunc = add
fmt.Println(f(2, 3))
```

---

## 9️⃣ Standard (Named) Function

A **named function** is defined normally.

```go
func greet() {
    fmt.Println("Hello")
}
```

---

## 🔟 Init Function

### 🔹 `init()`

Runs **before `main()`**, automatically.

```go
func init() {
    fmt.Println("Init called")
}
```

Rules:

* No parameters
* No return
* Multiple `init()` allowed
* Used for setup/config

---

## 1️⃣1️⃣ Function Expression

Assigning a function to a variable.

```go
var add = func(a, b int) int {
    return a + b
}
```

---

## 1️⃣2️⃣ Anonymous Function & IIFE

### Anonymous Function

```go
func(a int) {
    fmt.Println(a)
}(10)
```

### IIFE (Immediately Invoked Function Expression)

```go
func() {
    fmt.Println("IIFE")
}()
```

---

## 1️⃣3️⃣ A “Noob Function” (Beginner Mistake Example)

```go
func add(a int, b int) {
    a + b // ❌ no return, useless
}
```

Correct:

```go
func add(a, b int) int {
    return a + b
}
```

---

## 1️⃣4️⃣ Parameters vs Arguments

### Parameters

Defined in function signature:

```go
func add(a int, b int)
```

### Arguments

Actual values passed:

```go
add(2, 3)
```

---

## 1️⃣5️⃣ First Order Function

A function that **does not accept or return another function**.

```go
func square(x int) int {
    return x * x
}
```

---

## 1️⃣6️⃣ Higher Order Function

A function that **accepts or returns another function**.

```go
func operate(a int, b int, fn func(int, int) int) int {
    return fn(a, b)
}
```

---

## 1️⃣7️⃣ First Class Function

Functions in Go can be:

* Assigned to variables
* Passed as arguments
* Returned from functions

```go
func getAdder() func(int, int) int {
    return func(a, b int) int {
        return a + b
    }
}
```

---

## 1️⃣8️⃣ Callback Function

A function passed and executed later.

```go
func process(value int, callback func(int)) {
    callback(value)
}

func printValue(x int) {
    fmt.Println(x)
}

process(10, printValue)
```

---

## 1️⃣9️⃣ Go Internal Memory Model

### 🔹 Code Segment

* Stores compiled instructions

### 🔹 Data Segment (Global Memory)

* Global & package variables

### 🔹 Stack

* Function calls
* Local variables
* Stack frames

### 🔹 Heap

* Dynamically allocated memory
* Managed by GC

### 🔹 Garbage Collector (GC)

* Automatically frees unused heap memory
* No manual `free()` like C/C++

---

## 🔁 Stack Frame

Each function call creates a **stack frame**:

* Parameters
* Local variables
* Return address

Destroyed when function returns.

---

## 2️⃣0️⃣ Program Execution Phases

### 1️⃣ Compilation Phase

* Syntax check
* Type check
* Binary generation

### 2️⃣ Execution Phase

* `init()` runs
* `main()` runs
* Program executes

---

## 2️⃣1️⃣ go run vs go build

### `go run main.go`

```text
Compile → Create temp binary → Execute → Delete binary
```

### `go build main.go`

```text
Compile → Create ./main binary
```

Then:

```bash
./main
```
