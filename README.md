# Razorpay Intern Training - Golang Basics

## Golang Day 1

### Concepts Covered:

---

### Getting Started - Hello World
- Introduction to Golang and writing your first program.

---

### Packages
- [Tour of Go - Packages (Part 1)](https://tour.golang.org/basics/1)  
- [Tour of Go - Packages (Part 2)](https://tour.golang.org/basics/2)  
- [Tour of Go - Packages (Part 3)](https://tour.golang.org/basics/3)  

---

### Variables & Values
- [Tour of Go - Variables (Part 1)](https://tour.golang.org/basics/8)  
- [Tour of Go - Variables (Part 2)](https://tour.golang.org/basics/9)  
- [Tour of Go - Variables (Part 3)](https://tour.golang.org/basics/10)  
- [Go by Example - Values](https://gobyexample.com/values)  

---

### Flow Control Statements & Defer

#### For Loop
- [Tour of Go - For Loop (Part 1)](https://tour.golang.org/flowcontrol/1)  
- [Tour of Go - For Loop (Part 2)](https://tour.golang.org/flowcontrol/2)  

#### Go’s While Loop (For Loop Alternative!)
- [Tour of Go - While Loop (Part 1)](https://tour.golang.org/flowcontrol/3)  
- [Tour of Go - While Loop (Part 2)](https://tour.golang.org/flowcontrol/4)  

#### If-Else Statements
- [Tour of Go - If-Else (Part 1)](https://tour.golang.org/flowcontrol/5)  
- [Tour of Go - If-Else (Part 2)](https://tour.golang.org/flowcontrol/6)  
- [Tour of Go - If-Else (Part 3)](https://tour.golang.org/flowcontrol/7)  

#### Switch Statement
- [Tour of Go - Switch (Part 1)](https://tour.golang.org/flowcontrol/9)  
- [Tour of Go - Switch (Part 2)](https://tour.golang.org/flowcontrol/10)  
- [Tour of Go - Switch (Part 3)](https://tour.golang.org/flowcontrol/11)  

#### Defer Statement
- [Tour of Go - Defer (Part 1)](https://tour.golang.org/flowcontrol/12)  
- [Tour of Go - Defer (Part 2)](https://tour.golang.org/flowcontrol/13)  

---

### Functions
- [Tour of Go - Functions (Part 1)](https://tour.golang.org/basics/4)  
- [Tour of Go - Functions (Part 2)](https://tour.golang.org/basics/5)  
- [Tour of Go - Functions (Part 3)](https://tour.golang.org/basics/6)  
- [Tour of Go - Functions (Part 4)](https://tour.golang.org/basics/7)  

#### Additional Topics:
- [Variadic Functions](https://gobyexample.com/variadic-functions)  
- [Closures](https://gobyexample.com/closures)  
- [Recursion](https://gobyexample.com/recursion)  

---

### Structs, Slices, and Maps

#### Pointers
- [Go by Example - Pointers](https://gobyexample.com/pointers)  

#### Arrays
- [Go by Example - Arrays](https://gobyexample.com/arrays)  

#### Slices
- [Go by Example - Slices](https://gobyexample.com/slices)  

#### Maps
- [Go by Example - Maps](https://gobyexample.com/maps)  

#### Range
- [Go by Example - Range](https://gobyexample.com/range)  

#### Structs
- [Go by Example - Structs](https://gobyexample.com/structs)  
- [Tour of Go - Structs (Part 1)](https://tour.golang.org/moretypes/4)  
- [Tour of Go - Structs (Part 2)](https://tour.golang.org/moretypes/5)  

---

### Methods and Interfaces

#### Methods
- [Tour of Go - Methods](https://tour.golang.org/methods/1) to [Methods (Part 8)](https://tour.golang.org/methods/8)  

#### Interfaces
- [Tour of Go - Interfaces](https://tour.golang.org/methods/9) to [Interfaces (Part 14)](https://tour.golang.org/methods/14)  

---

### Error Handling
- [Go by Example - Errors](https://gobyexample.com/errors)






## Data Race Prevention:
One of the key benefits of Go’s concurrency model is the prevention of data races. Data races occur when multiple goroutines access and modify shared data without proper synchronization. In Go, data races are discouraged by design.

Goroutines and channels promote safe concurrent programming. Since goroutines execute concurrently but not necessarily simultaneously, you can avoid race conditions by ensuring that only one goroutine can access a piece of shared data at a time through proper synchronization with channels or other synchronization primitives.