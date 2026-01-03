# 🚀 Phase 1 – Lesson 2

## Functions in Go (Deep Dive, Production-Level)

By the end of this lesson, you will:

* Fully understand **Go functions**
* Master **multiple return values**
* Understand **error-first design**
* Learn **pass by value vs reference**
* Think like a **professional Go engineer**, not a JS dev

---

## 1️⃣ What Is a Function in Go?

A function is a **named block of reusable code**.

### Basic function

```go
func add(a int, b int) int {
    return a + b
}
```

### Calling a function

```go
result := add(2, 3)
fmt.Println(result) // 5
```

---

## 2️⃣ Function Syntax (Rules You Must Know)

```go
func functionName(paramName type, paramName type) returnType {
    // code
}
```

Key rules:

* Parameter **name first**, type second
* Return type comes **after parameters**
* Curly braces are mandatory
* No arrow functions ❌

---

## 3️⃣ Multiple Parameters with Same Type

Instead of:

```go
func add(a int, b int) int
```

You can write:

```go
func add(a, b int) int
```

Very common in Go code.

---

## 4️⃣ Multiple Return Values (🔥 Go Signature Feature)

Go functions can return **multiple values**.

### Example

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### Calling it

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

💡 This replaces exceptions.

---

## 5️⃣ Error-First Return Pattern (CRITICAL)

Convention:

```go
(value, error)
```

Example:

```go
data, err := fetchData()
if err != nil {
    return err
}
```

Why Go does this:

* Errors are **values**
* No hidden control flow
* Easy to reason about
* Safer in large systems

---

## 6️⃣ Ignoring Return Values (`_` blank identifier)

If you don’t need a return value:

```go
_, err := divide(10, 2)
```

Or:

```go
value, _ := divide(10, 2)
```

💡 `_` means “I intentionally ignore this”.

---

## 7️⃣ Named Return Values

You can **name return variables**.

```go
func sum(a, b int) (result int) {
    result = a + b
    return
}
```

⚠️ Use sparingly — clarity first.

---

## 8️⃣ Early Returns (Go Style)

Go prefers **early return**, not deep nesting.

❌ Bad:

```go
if err == nil {
    if ok {
        return value
    }
}
```

✅ Good:

```go
if err != nil {
    return err
}
if !ok {
    return nil
}
return value
```

This is **idiomatic Go**.

---

## 9️⃣ Pass by Value vs Pass by Reference (VERY IMPORTANT)

### Go is ALWAYS pass by value

```go
func change(x int) {
    x = 10
}

func main() {
    a := 5
    change(a)
    fmt.Println(a) // 5
}
```

---

### Using pointers to modify data

```go
func change(x *int) {
    *x = 10
}

func main() {
    a := 5
    change(&a)
    fmt.Println(a) // 10
}
```

💡 Go gives you **explicit control** over mutation.

---

## 🔁 JavaScript Comparison

### JavaScript:

* Objects passed by reference (sort of)
* Primitive vs reference confusion
* Implicit mutation

### Go:

* Everything passed by value
* Pointers are explicit
* Much safer mental model

---

## 10️⃣ Functions Are First-Class Citizens

Functions can be:

* Assigned to variables
* Passed as arguments
* Returned from functions

### Function as variable

```go
add := func(a, b int) int {
    return a + b
}
```

### Passing function as argument

```go
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}
```

---

## 11️⃣ Anonymous Functions

```go
func() {
    fmt.Println("Hello")
}()
```

Used for:

* Callbacks
* Goroutines
* Closures

---

## 12️⃣ Closures (Like JS, But Safer)

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

---

## 🧪 Mini Practice (IMPORTANT)

Predict output:

```go
func modify(x int, y *int) {
    x = 20
    *y = 30
}

func main() {
    a := 10
    b := 10

    modify(a, &b)

    fmt.Println(a, b)
}
```

---

## 🎯 Lesson 2 Outcome

You now understand:

* Function syntax
* Multiple return values
* Error-first pattern
* Pointers in functions
* Closures & first-class functions

This is **core Go mastery**.

---

## ▶️ Next Lesson

**Phase 1 – Lesson 3**:

> **“Phase 1 – Lesson 3”**
