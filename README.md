# Belajar Golang Dasar

A basic Go programming learning project to explore fundamental concepts and features of the Go programming language.

## Project Overview

This project serves as a foundation for learning Go (Golang) programming language basics. It contains simple programs and examples to help understand core Go concepts.

## Project Structure

```
belajar-golang-dasar/
├── go.mod          # Go module definition
├── helloworld.go   # Basic Go program template
└── README.md       # This file
```

## Requirements

- Go 1.24.4 or later
- Basic understanding of programming concepts

## Getting Started

### Installation

1. Make sure you have Go installed on your system. You can download it from [golang.org](https://golang.org/download/)

2. Clone or download this project to your local machine

3. Navigate to the project directory:
   ```bash
   cd belajar-golang-dasar
   ```

### Running the Program

To run the Hello World program:

```bash
go run helloworld.go
```

To build an executable:

```bash
go build helloworld.go
```

### Module Management

This project uses Go modules. To initialize or update dependencies:

```bash
# Initialize module (already done)
go mod init belajar-golang-dasar

# Download dependencies
go mod download

# Clean up unused dependencies
go mod tidy
```

## Learning Objectives

This project aims to cover the following Go fundamentals:

- [ ] Basic syntax and program structure
- [ ] Variables and data types
- [ ] Functions and methods
- [ ] Control structures (if/else, loops)
- [ ] Arrays, slices, and maps
- [ ] Structs and interfaces
- [ ] Error handling
- [ ] Concurrency with goroutines
- [ ] Packages and modules

## Development

### Code Style

Follow standard Go formatting conventions:

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...
```

### Testing

Add tests for your Go programs:

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)

## License

This project is for educational purposes only.

## Contributing

Feel free to add more examples and improve the code as you learn Go programming concepts.
