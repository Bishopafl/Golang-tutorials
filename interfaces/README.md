# Understanding Interfaces in Go

## 🌐 What is an Interface?
In Go, an **interface** is a type that defines a set of method signatures.  
Any type that implements those methods **automatically satisfies the interface** — no explicit `implements` keyword is required (unlike Java or C#).

Think of an interface as a **contract**:  
"If your type can do these things (methods), then it is this interface."

---

## ✨ Why Interfaces Matter
- **Decoupling**: They allow code to depend on *behavior* rather than *concrete implementations*.
- **Polymorphism**: Different types can be treated the same if they implement the same interface.
- **Testability**: Makes mocking and substituting dependencies easier.
- **Flexibility**: Functions can work with any type that satisfies an interface.

---

## 📦 Basic Example

### Defining an Interface
```go
package main

import "fmt"

// Speaker defines a behavior (contract).
type Speaker interface {
	Speak() string
}
```

### Implementing the Interface
```go
// Dog implicitly implements Speaker
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat implicitly implements Speaker
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}
```

### Using the Interface
```go
func main() {
	var s Speaker

	s = Dog{}
	fmt.Println(s.Speak()) // Output: Woof!

	s = Cat{}
	fmt.Println(s.Speak()) // Output: Meow!
}
```

🔑 Key Points

1. Implicit Satisfaction
    - No implements keyword — if a type has the required methods, it satisfies the interface.
2. Empty Interface (interface{})
    - Special case that accepts any type.
    - Often replaced by any (alias in Go 1.18+).
3. Interfaces as Contracts
    - Interfaces should be small and focused.
    - Prefer many small interfaces over one large one.
    - Example: io.Reader, io.Writer — both define a single method.