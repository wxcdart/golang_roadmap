# Go Practical Exercises: Getting Started

This guide provides hands-on prompts to help you master the basics of Go. Work through each exercise, read the docs as needed, and experiment!

---

## 1. Hello World & CLI Basics
- Write a Go program that prints "Hello, World!" to the terminal.
- Modify it to accept your name as a command-line flag (e.g., `-name=Alice`) and greet you by name.
- Add a second flag (e.g., `-shout`) that, if set, prints the greeting in uppercase.

## 2. Simple File I/O
- Write a program that creates a text file and writes a few lines to it.
- Extend it to read the file back and print its contents to the terminal.
- Add error handling for file operations (open, read, write).

## 3. Parsing Flags
- Use the `flag` package to parse multiple command-line flags (string, int, bool).
- Write a program that takes a filename and a number N as flags, writes N lines to the file, then reads and prints them.

## 4. Basic Data Structures
- Implement a stack using a slice of integers. Add push, pop, and print methods.
- Create a map that counts word frequencies in a string. Print the map.
- Define a `Person` struct with fields for name and age. Write a function that takes a slice of `Person` and prints the oldest person.

---

**Tips:**
- Use `go run` to quickly test your code.
- Read the official docs for [`flag`](https://pkg.go.dev/flag), [`os`](https://pkg.go.dev/os), [`io`](https://pkg.go.dev/io), and [`fmt`](https://pkg.go.dev/fmt).
- Try to write small, self-contained programs for each exercise.

*Save your solutions and revisit them as you learn more!*
