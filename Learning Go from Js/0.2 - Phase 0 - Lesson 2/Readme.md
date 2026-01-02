# 🧠 Phase 0 – Lesson 2

## Go Installation, Tooling & Your First Program

By the end of this lesson, you will:

* Understand **how Go runs**
* Be comfortable with the **Go toolchain**
* Write and run your **first real Go program**
* Know how **modules** work (VERY IMPORTANT)

---

## 1️⃣ Installing Go (Correct Way)

### Check if Go is installed

```bash
go version
```

Expected output (example):

```
go version go1.22.0 darwin/amd64
```

If not installed:

* Download from **go.dev**
* Install normally (no custom setup needed)

💡 Go installs everything in one step. No Node/NVM-like chaos.

---

## 2️⃣ Understanding Go Environment (`go env`)

Run:

```bash
go env
```

Important variables (don’t memorize yet):

* `GOROOT` → Go installation
* `GOPATH` → workspace (legacy, still used)
* `GOOS`, `GOARCH` → target OS & architecture
* `GOMOD` → active module

🧠 **Modern Go uses modules**, not GOPATH projects.

---

## 3️⃣ How Go Code Is Organized

### Every Go file belongs to a package

```go
package main
```

### Two types of packages:

* `main` → executable
* others → libraries

Only `main` has:

```go
func main()
```

---

## 4️⃣ Your First Go Program (Step-by-Step)

### Step 1: Create a project folder

```bash
mkdir hello-go
cd hello-go
```

### Step 2: Initialize a Go module

```bash
go mod init hello-go
```

This creates:

```
go.mod
```

📄 `go.mod` tells Go:

* module name
* Go version
* dependencies

➡️ Think of `go.mod` as `package.json`, but **simpler and stricter**.

---

### Step 3: Create `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

---

### Step 4: Run the program

```bash
go run main.go
```

Output:

```
Hello, Go
```

🎉 You just ran your first Go program.

---

## 5️⃣ `go run` vs `go build`

### `go run`

```bash
go run main.go
```

* Compiles
* Executes
* Deletes binary afterward
* Good for development

### `go build`

```bash
go build
```

* Creates executable binary
* No execution
* Binary name = module or folder name

Run it:

```bash
./hello-go
```

🔥 This binary can be deployed **without Go installed**.

---

## 6️⃣ Go Compilation Model (Very Important)

When you run:

```bash
go run main.go
```

Go does:

1. Compile source code
2. Link dependencies
3. Create binary
4. Execute binary

No interpreter.
No runtime dependency.

---

## 7️⃣ Imports (How Go Manages Dependencies)

```go
import "fmt"
```

* Standard library first
* Explicit imports only
* No unused imports allowed ❌

If you do:

```go
import "fmt"
// but don't use fmt
```

➡️ **Compilation fails**

This enforces clean code.

---

## 8️⃣ Formatting (Go Is Strict)

Run:

```bash
gofmt -w main.go
```

* Enforces a single code style
* No debates, no bikeshedding
* Industry standard

Most editors auto-run this.

---

## 9️⃣ Go File Naming Rules

* `main.go`, `server.go`, `handler.go`
* Lowercase
* No spaces
* One package per folder

---

## 10️⃣ JavaScript vs Go (Quick Comparison)

| Concept   | JavaScript   | Go            |
| --------- | ------------ | ------------- |
| Runtime   | Node.js      | Native binary |
| Types     | Dynamic      | Static        |
| Async     | Event loop   | Goroutines    |
| Modules   | package.json | go.mod        |
| Formatter | Optional     | Mandatory     |

---

## 🎯 Lesson 2 Outcome

You now understand:

* How Go is installed
* How Go code is structured
* `go run` vs `go build`
* Modules (`go mod`)
* Your first Go program

You are officially **inside the Go ecosystem**.

---

## ▶️ Next Lesson

> **“Phase 0 – Lesson 3”**
