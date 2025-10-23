# Learning Go programming

---

> **Official website**: [**go.dev**](https://go.dev/)

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

# Packages
Go code must be organised into packages:

- The package a program belongs to is stated using the `package <your_package_name>`
- "Package" here is an organisation concept (like a label), not a directory

Any Go project must have at least one package.

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

# `main` package as crucial to a module
However, `main` as a package is a reserved package name within Go's code organisation:

- It marks the entrypoint for your application
- Other program files may belong to other packages, but at least one must belong to `main`
- Without a `main` package, no executable can be created for your application <br> I.e.: *Running `go build` within the module directory yields nothing without a `main` package*