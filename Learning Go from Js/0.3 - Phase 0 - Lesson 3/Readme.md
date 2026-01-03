# 🧠 Phase 0 – Lesson 3

## Go Syntax Basics: Variables, Types & Zero Values

By the end of this lesson, you will:

* Understand **Go’s type system**
* Know how variables really work
* Avoid JS-style runtime bugs
* Think in **compile-time safety**

---

## 1️⃣ Go Is Strongly & Statically Typed

### JavaScript:

```js
let x = 10
x = "hello" // allowed
```

### Go:

```go
var x int = 10
x = "hello" // ❌ compile-time error
```

💡 Go catches mistakes **before the program runs**.

---

## 2️⃣ Variable Declarations in Go

### 2.1 `var` keyword

```go
var age int = 25
```

You can also omit the type:

```go
var age = 25
```

Go infers the type **once** — then it’s fixed.

---

### 2.2 Short declaration operator `:=` (VERY COMMON)

```go
age := 25
```

Rules:

* Only works **inside functions**
* Cannot redeclare existing variable in same scope
* Go infers type automatically

💡 This is Go’s clean alternative to `let`.

---

### ❌ Invalid usage

```go
age := 25
age := 30 // ❌ redeclaration error
```

---

## 3️⃣ Multiple Variable Declaration

```go
var x, y int = 10, 20
```

or

```go
a, b := 1, 2
```

---

## 4️⃣ Constants (`const`)

```go
const pi = 3.14159
```

Rules:

* Immutable
* Must be known at compile time
* No `:=`

```go
const x := 10 // ❌ invalid
```

---

## 5️⃣ Basic Data Types in Go

### Numeric Types

```go
int, int8, int16, int32, int64
uint, uint8, ...
float32, float64
```

Default:

```go
x := 10      // int
y := 3.14    // float64
```

💡 Go is explicit — no auto casting.

---

### Boolean

```go
var isActive bool = true
```

No truthy / falsy like JS.

```go
if isActive { } // ✅
if 1 { }        // ❌
```

---

### Strings

```go
name := "Khalid"
```

Strings are:

* Immutable
* UTF-8 encoded

No `char` type — use `rune`.

---

## 6️⃣ Zero Values (VERY IMPORTANT)

In Go, **every variable has a default value**.

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| float   | 0.0        |
| bool    | false      |
| string  | ""         |
| pointer | nil        |
| slice   | nil        |
| map     | nil        |

Example:

```go
var x int
fmt.Println(x) // 0
```

💡 No `undefined` or `null` confusion like JS.

---

## 7️⃣ Type Conversion (Explicit Only)

### JavaScript:

```js
Number("10") + 1
```

### Go:

```go
x := 10
y := int(3.5) // y = 3
```

No implicit conversions.

```go
var a int = 10
var b float64 = a // ❌ error
```

Must convert explicitly:

```go
var b float64 = float64(a)
```

---

## 8️⃣ Scope Rules (Simple & Strict)

### Block scope:

```go
if true {
    x := 10
}
// x is not accessible here
```

### Function scope:

```go
func demo() {
    y := 20
}
```

No hoisting like JS.

---

## 9️⃣ Shadowing (Dangerous but Common)

```go
x := 10

if true {
    x := 20 // new variable, shadows outer x
}
```

💡 This is legal — but dangerous if misused.

---

## 10️⃣ Go Style Rules (Professional Standard)

* Use `camelCase`
* Exported names start with **capital letter**
* Unexported names start with **lowercase**
* Keep variables short but meaningful

```go
count := 10
userID := 5
```

---

## 🎯 Lesson 3 Outcome

You now understand:

* How Go variables work
* Static typing vs JS dynamic typing
* Zero values
* Explicit conversions
* Scope & shadowing

This is the **foundation of all Go code**.

---

## 🧪 Mini Practice (Do This Mentally or in Code)

Try to predict output:

```go
var x int
y := 5

if y > 3 {
    x := 10
    fmt.Println(x)
}

fmt.Println(x)
```

(We’ll discuss answers next lesson.)

---

## ▶️ Next Lesson

> **“Phase 1 – Lesson 1”**
