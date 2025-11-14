# Golang Tutorials

Welcome to my collection of **Go (Golang)** learning exercises and mini-projects! 🧠💻  
This repository is a growing collection of small tutorials, experiments, and challenges I've been completing as part of my ongoing journey to learn and master Go.

Each folder in this repo represents a focused concept or hands-on exercise, often inspired by Go courses, documentation, or personal curiosity.  
The goal is simple: **learn by building**, and share that learning process openly.

---

## 📂 Repository Structure
````markdown
| Folder | Description |
|--------|--------------|
| **cards/** | A simple program that demonstrates working with slices, loops, and randomization — building a basic card deck app in Go. |
| **channels/** | Examples of using Go channels to handle concurrency safely, including examples of sending and receiving data using `range` and `select`. |
| **even-odd/** | Small exercise that determines whether numbers are even or odd using Go's control flow and functions. |
| **file-to-term/** | Reads a file and prints its contents to the terminal. Demonstrates file I/O operations using the `os` and `io` packages. |
| **hello-world/** | The classic “Hello, World!” — my first step into Go! Covers Go’s basic syntax, `main()` function, and package structure. |
| **http/** | Introduces the `net/http` package and how to create custom `io.Writer` implementations (like a `logWriter`) to handle HTTP responses. |
| **interfaces/** | Shows how Go interfaces work by implementing simple examples involving the `net/http` package and struct-based receivers. |
| **map/** | Demonstrates the creation, initialization, and usage of maps in Go, including how to iterate over key-value pairs. |
| **shapes/** | Object-oriented style Go tutorial demonstrating interfaces and struct implementations for geometric shapes (e.g., `square`, `triangle`). |
| **structs/** | Explores Go structs, pointers, and receiver functions — focusing on updating data inside structs through methods. |
| **tests/** | Basic Go test examples following Go’s built-in testing package. Used to test logic from various exercises (e.g., quiz app challenges). |
````
---

## 🧩 Goals

- Build a solid foundation in Go syntax and idioms.
- Learn core Go concepts: interfaces, structs, maps, concurrency, channels, and testing.
- Practice writing clean, idiomatic Go code.
- Document progress and share examples for others learning Go.

---

## ⚙️ How to Run

To run any of the tutorials:

```bash
# Navigate to a specific tutorial
cd channels

# Run the Go program
go run main.go
````

Some projects may include additional files (like `deck.go` or `helpers.go`) that are imported automatically when using `go run .`.

---

## 💬 Notes

These projects are **learning experiments**, not production-ready applications.
If you’re also learning Go, feel free to fork the repo, explore, and even share suggestions or improvements!

---

## 🧠 Learning Resources

Here are some of the resources I’ve been following:

* [Go by Example](https://gobyexample.com/)
* [The Go Programming Language Tour](https://tour.golang.org/)
* [Udemy Go Courses by Stephen Grider](https://www.udemy.com/course/go-the-complete-developers-guide/)
* [Go Docs](https://pkg.go.dev/std)

---

### 🧭 About This Repo

Each commit shows my incremental learning progress.
Some folders may be refactored or rewritten over time as I gain a deeper understanding of Go concepts.

> *“The best way to learn Go is to build with it.”* — That’s exactly what I’m doing here.

---

**Author:** Adam Lopez
**Language:** Go (1.x+)
**License:** MIT

```
