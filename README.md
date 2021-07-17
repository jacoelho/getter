# Getter

Go struct field getter code generator.

## Usage

Introduce `go:generate` instruction.

```go
//go:generate getter -type Foo,Bar
```

## Motivation

Go does not requires getters and is usally considered idiomatic to access public fields directly.

```go
type Example struct {
    Name string
}

var e Example
e.Name = "some example"

fmt.Println(e.Name)
```

However, if the field, for example, is a pointer you may have a nil deference error if it is not initialised:

```go
type ExampleHeader struct {
	Name string
	File string
}

type Example struct {
	Header *ExampleHeader
}

var e Example
fmt.Println(e.Header.Name)

// panic: runtime error: invalid memory address or nil pointer dereference
```

The simplest approach is verify if it is `nil` before access:

```go
if e.Header!= nil {
    fmt.Println(e.Header.Name)
}
```

Although it is an idiomatic solution, it gets old really fast and prune to error if you are working with multiple nested structures. The end solution may end-up with something like:

```go
if e.Header != nil && e.Header.Name != nil && e.Header.Name.Title != nil {
    ...
}
```

Getter generates code that allows chaining gets without fearing a nil deference:

```go
type Name struct {
	Title  string
	Author string
}

type Header struct {
	Name *Name
}

type Example struct {
	Example *Header
}

e.GetExample().GetName().GetTitle()

// ""
```

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.