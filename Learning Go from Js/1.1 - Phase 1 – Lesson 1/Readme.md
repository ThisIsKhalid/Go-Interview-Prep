# 🚀 Phase 1 – Lesson 1

## Control Flow in Go (if, for, switch, defer)

This lesson will **rewire your JavaScript habits** into **idiomatic Go thinking**.

By the end of this lesson, you’ll:

* Fully understand Go control flow
* Stop searching for `while` loops 😄
* Write clean, readable logic
* Use `defer` like a professional Go developer

---

## 1️⃣ `if` Statement (Go Style)

### Basic `if`

```go
age := 20

if age >= 18 {
    fmt.Println("Adult")
}
```

Rules:

* Condition **must be boolean**
* Parentheses are **not allowed**

❌ Invalid:

```go
if (age >= 18) { }
```

---

### `if` with `else`

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

---

### `if` with initialization (VERY IMPORTANT)

```go
if score := 85; score >= 50 {
    fmt.Println("Passed")
} else {
    fmt.Println("Failed")
}
```

💡 `score` exists **only inside this if block**.

This pattern is used heavily in **error handling**:

```go
if err := doSomething(); err != nil {
    return err
}
```

---

## 2️⃣ `for` Loop (The ONLY Loop in Go)

Go has **no**:

* `while`
* `do-while`
* `forEach`

Everything is `for`.

---

### Classic `for`

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

### `while`-style loop

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

---

### Infinite loop

```go
for {
    fmt.Println("Running...")
}
```

(Used in servers, workers, listeners)

---

### Loop with `break` & `continue`

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

---

## 3️⃣ `switch` (Much More Powerful Than JS)

### Basic switch

```go
day := 2

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

💡 No `break` needed — Go breaks automatically.

---

### Multiple cases

```go
switch day {
case 1, 2, 3:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
}
```

---

### Switch without expression (🔥 VERY IMPORTANT)

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

This replaces long `if-else` chains cleanly.

---

## 4️⃣ `defer` (Go Superpower)

`defer` means:
👉 **Run this when the function returns**

### Basic example

```go
func demo() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

Output:

```
Hello
World
```

---

### Common real-world use cases

#### 1. Closing files

```go
file, err := os.Open("data.txt")
if err != nil {
    return err
}
defer file.Close()
```

#### 2. Unlocking mutex

```go
mutex.Lock()
defer mutex.Unlock()
```

💡 This prevents **resource leaks**.

---

### Multiple `defer` calls (LIFO order)

```go
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
```

Output:

```
3
2
1
```

(Stack behavior)

---

## 5️⃣ JavaScript vs Go (Control Flow)

| Concept   | JavaScript          | Go            |
| --------- | ------------------- | ------------- |
| Loops     | for, while, forEach | only `for`    |
| Condition | truthy/falsy        | strictly bool |
| switch    | needs break         | auto break    |
| finally   | try/finally         | `defer`       |

---

## 6️⃣ Common Beginner Mistakes (Avoid These)

❌ Using parentheses in `if`
❌ Expecting truthy/falsy
❌ Forgetting scope of variables
❌ Not using `defer` for cleanup

---

## 🧪 Mini Practice (Important)

Predict output:

```go
func test() {
    defer fmt.Println("A")
    defer fmt.Println("B")

    for i := 0; i < 2; i++ {
        defer fmt.Println(i)
    }
}
```

Call:

```go
test()
```

👉 Write the output order mentally.

---

## 🎯 Lesson 1 Outcome

You now:

* Control program flow in Go
* Understand Go’s looping philosophy
* Use `switch` properly
* Use `defer` safely and professionally

This is **real Go**, not JS-in-Go-syntax.

---

## ▶️ Next Lesson

**Phase 1 – Lesson 2**:

> **“Phase 1 – Lesson 2”**
