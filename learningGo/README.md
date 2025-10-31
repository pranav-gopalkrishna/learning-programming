<h1>Learning Go programming</h1>

---

> **Official website**: [**go.dev**](https://go.dev/)

---

**Contents**:

- [Why Go?](#why-go)- [Why Go?](#why-go)
- [How to compile and run a Go program?](#how-to-compile-and-run-a-go-program)
- [Module](#module)
- [Packages](#packages)
  - [`main` package as crucial to a module](#main-package-as-crucial-to-a-module)
  - [`main` function as crucial to an executable module](#main-function-as-crucial-to-an-executable-module)
- [Variables and types](#variables-and-types)
- [Variable type conversion and type assignment](#variable-type-conversion-and-type-assignment)
  - [Zero values](#zero-values)
  - [Walrus operator](#walrus-operator)
- [Pointers](#pointers)
- [Functions](#functions)
- [Error handling](#error-handling)
- [User-defined types using `struct`](#user-defined-types-using-struct)

---

# Why Go?
- Focus on simplicity and clarity (i.e. readability) (like Python)
- Focus on performance and concurrency (like C and C++)
- Many built-in core features that come as a standard library
- Type safety

# How to compile and run a Go program?
```
go run hello_world.go
```

You can replace `hello_world.go` with the name of your program source file.

# Module
Module = Go project (can contain one or more packages)

A directory is marked as a module using the `go.mod` file; this can be created as follows:

- Navigate to the directory you want to mark as a module
- Run `go mod init xyz`
    - Here, `xyz` is your module name
    - It can be any string
- The `go.mod` file contains:
    - The chosen name for the module
    - The Go version

An example of a `go.mod` file:

```
module oh_wow_a_module

go 1.25.3
```

# Packages
Go code must be organised into packages:

- The package a program belongs to is stated using the `package <your_package_name>`
- "Package" here is an organisation concept (like a label), not a directory

Any Go project must have at least one package.

---

Across files belonging to the same packages:

- Function names and global variables are shared
- Imports are not shared

---

Working with multiple packages:

- Multiple packages cannot be defined in a single workspace (i.e. directory)
- Each package must have its own directory or subdirectory
- A package can be split across multiple directories, but one directory cannot have files belonging to multiple packages <br> **NOTE**: *But one directory can have subdirectories, where files may belong to other packages*
- For executable code:
    - The `main` package should be in the parent directory (i.e. the module's directory)
    - Other packages should be in the same directory but in subdirectories
- Only variables and functions with names starting with an uppercase are exported <br> I.e.: *Only such variables are accessible to other packages via imports* <br> **NOTE**: *This is not a requirement for variables or functions within the same package, even if in different files*

An example for showing the basics of multi-package modules:

[`my_first_multi_package_module`](./my_first_multi_package_module/)

---

Working with 3rd party packages:

- Within the module's workspace, run `go get <path of the package>` <br> E.g.: `go get github.com/Pallinder/go-randomdata`
- This package is stored globally in your system
- A line gets added to the `go.mod` file of the module: <br> `require <path of the package>`
- To import this package:
    - Use `import` for the path as given above
    - Use the package name exposed to refer to it in the code

## `main` package as crucial to a module
However, `main` as a package is a reserved package name within Go's code organisation:

- It marks the entrypoint for your application
- Other program files may belong to other packages, but at least one must belong to `main`
- Without a `main` package, no executable can be created for your application <br> I.e.: *Running `go build` within the module directory yields nothing without a `main` package*

## `main` function as crucial to an executable module
- `main` function is the entrypoint of executing for your application
- Without it, the source code will compile but nothing will execute
- There should be only 1 `main` function in the whole module
- The `main` function should be defined in exactly 1 program belonging to the `main` package

**NOTE**: *`main` is not necessary for a library, only an executable module.*

# Exporting programmatic objects
**NOTE**: *Programmatic Object = Variable, Constant, Function, User-Defined Type, etc.*

- Starting with uppercase => Is exported to another package if imported in the other package
- Starting with lowercase => Is not exportable to another package; for in-package use only

**NOTE**: This applies for:

- User-defined type names and fields (see: [User-defined types using `struct`](#user-defined-types-using-struct)) <br> **NOTE**: *Even field names should start with uppercase to be exportable, i.e. accessible by the type when the type is exported*
- Methods for user-defined types
- Functions
- Variables

**NOTE**: *For in-package use, both uppercase and lowercase starting names work.*

# Variables and types
# Variable type conversion and type assignment
Go is a type-safe language, and unlike some other programming languages, there is no implicit type conversion, even between integer and float types. Hence, you must not only define the type of the variable, but you must also typecast variables where necessary to ensure type compatibility. For example:

```go
var a = 0 // Is of type int
var b = 0.2 // Is of type float64
```

Here we have that:

- `var c = a + b` would lead to a type incompatibility error
- `var c = float64(a) + b` needs to be done to ensure `c` is of `float64` type

**NOTE**: *Typecasting functions are built-in in Go.*

---

Crucially, we can specify the variable type explicitly as follows:

```go
var a float64 = 0 // Is of type float64
```

This way, we do not have to:

- Depend on Go to infer the type
- Put unnecessary decimal points to ensure the type is inferred as `float64`

## Zero values
All Go value types come with a so-called "zero value" which is the value stored in a variable if no other value is explicitly set. For example, the following int variable would have a default value of 0 (because 0 is the null value of `int`, `int32`,  etc):

```go
var age int // age is 0
```

Here's a list of the zero values for the different types:

- `int` => 0
- `float64` => 0.0
- `string` => "" (i.e. an empty string)
- `bool` => false
- Any pointer => `nil` (a special value built into Go)

## Walrus operator
The := operator in Go provides a concise way to declare and initialise variables, with Go inferring the variable's type based on the assigned value. Note that this operator must only be used for declaring and initialising variables, not for assigning a declared variable with a new value.

# Pointers
Pointer = Variable that stores a memory address instead of a value

Why use pointers:

- Avoid duplication of data (Go functions create copies of a variable when passing them as arguments to functions) <br> **NOTE**: *This is relevant for variables containing large amounts of data*
- Direct mutation of variable value (instead of having to assign value to the variable later)

Note that:

- Pointers are type specific (integer pointer, float pointer, etc.)
- A pointer type is specified as `*<type>`, e.g.: `*int` (integer pointer)\
- The memory address of a variable is obtained by `&` <br> E.g.: *For variable*, `b := 2`, `&b` *gives the integer pointer value of this variable*
- To read/obtain the value referred to by a pointer `p`, use `*`, e.g.: `*p`

**NOTE**: *Type matters because memory space allocation depends on the type; hence reading using a pointer not only means knowing the address of the first byte but the range of bytes containing the desired value. Hence, specifying pointer type is essential. Note that the memory address of a variable, as obtained by the `&` operator, is automatically cast in the appropriate pointer type.*


Zero value:

- Zero value of any pointer is `nil` (a special value built into Go)
- `nil` represents the absence of an address value, i.e. a pointer pointing at no address/no value in memory

# Functions
Some key points/tips:

- A method is a function that is accessible via a type instance (i.e. a variable of a certain type)
- Arguments passed to the function are duplicates of the actual variables passed
- Returns of a function are duplicates of the actual variables used
- When all the arguments of a function are the same type, specifying the type for the last argument is sufficient

# Error handling
- Error data is encapsulated by the `error` type, whose zero-value is `nil` <br> **NOTE**: *An error value being `nil` means no error*
- By default, Go returns errors as data, rather than stopping program execution
- The built-in `errors` package contains methods to initialise `error` variables

# User-defined types using `struct`
## Introduction
A type is an interface through which data can be handled; here, interface implies:

- Data fields
- Arrangement of data

In Go, a `struct` (short for "structure"):

- Is a user-defined type that groups together 0 or more named fields of 1 or more types into a single unit
- Provides a way to create complex data-handling interfaces by encapsulating related data

## Key points
- When a user-defined type variable is passed to a function/method its duplicate is used <br> **NOTE**: *This is the case with any variable passed to a function/method*
- A method can be defined for a type by defining a receiver argument (or simply receiver) as the desired type <br> E.g.: *Instead of* `func f(arg mytype) {...}`, *you do* `func (arg mytype) f() {...}` <br> **NOTE**: *The first key point applies for receiver arguments as well*
- A mutation method (which changes the original variable) can be defined as follows: <br> *Instead of* `func f(arg *mytype) {...}`, *you do* `func (arg *mytype) f() {...}` <br> **NOTE** (given a variable `myvar` of type `mytype`):
    - The Go compiler ensures that doing `myvar.f` passes the pointer of `myvar` or its duplicate, however defined
    - Like pointers, struct-defined types do not need a dereference operator to access the field values <br> I.e.: *If* `myvarp = &myvar`, *then* `myvar.field1` *and* `myvarp.field1` *both obtain the value of* `.field1` <br> (However, using the dereference operator `((*myvarp).field1` is also valid, though unnecessary)


## Initialising a user-defined type variable
For example, take the following user-defined type:

```go
type mytype struct {
    field1 string
    field2 string
}
```

E.g.: Assigning values to each field by mentioning field names:

```go
var x = mytype {
    field1: "abc"
    field2: "123"
}
```

E.g.: Assigning values to select fields by mentioning field names:

```go
var x = mytype {
    field1: "abc"
}
```

*In the above example, `field2` is assigned the appropriate zero-value ("" in this case).*

E.g.: Assigning values to to all fields by giving values in defined order:

```go
var x = mytype {
    "abc", // field1
    "123"  // field2
}
```

## Creation/constructor function for a user-defined type
- No in-built programmatic structure to do so
- The convention is to name such a function `new...` (`...` being the type name or something)

For example, take the following user-defined type:

```go
type mytype struct {
    field1 string
    field2 string
}
```

E.g.: Constructor approach 1:

```go
func newMyType(field1, field2 string) mytype {
    return mytype {
        field1: field1
        field2: field2
    }
}
```

E.g.: Constructor approach 2:

```go
func newMyType(field1, field2 string) *mytype {
    return &mytype {
        field1: field1
        field2: field2
    }
}
```

Differences lie in the fact that:

- Approach 1 duplicates an initialised variable and returns it <br> Hence: *Assign it to a variable of type `mytype`*
- Approach 2 sends the initialised variable's reference <br> Hence: *Assign it to a pointer of type `mytype`, i.e. a variable of type `*mytype`*

Interestingly, fields and methods of a struct variable can be accessed using the same interface, whether you are using the struct variable itself or a pointer of the corresponding type. In other words, whether `x` is a variable of type `mytype` or type `*mytype`, `x.field1` gives you the value of `field1`.

## Exporting a user-defined type and its fields
See: [Exporting programmatic objects](#exporting-programmatic-objects)

**NOTE**:

- Exporting the user-defined type alone does not ensure its fields are exposed to exports
- To expose fields for exports, you must ensure they too start with uppercase
- It may not always be relevant to expose the fields <br> E.g.: *If fields are to be modified purely via methods*
- To expose methods for exports, you must ensure they too start with uppercase <br> **NOTE**: *Naturally, this includes the constructor method too*

# User-defined type aliases
- Existing types can be aliased and defined as a separate type using `type`
- This allows defining specific methods for a certain purpose

